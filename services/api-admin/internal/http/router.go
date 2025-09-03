package httpserver

import (
	"context"
	"net/http"
	"time"
	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Router devolve um http.Handler pronto pra usar no http.ListenAndServe.
// Ajuste o tipo abaixo se seu store.Open não retorna *pgxpool.Pool.
func Router(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // ajuste conforme necessário
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Healthchecks
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
	})

	r.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		if db != nil {
			if err := db.Ping(ctx); err != nil {
				writeJSON(w, http.StatusServiceUnavailable, map[string]any{
					"status": "degraded",
					"error":  err.Error(),
				})
				return
			}
		}
		writeJSON(w, http.StatusOK, map[string]any{"status": "ready"})
	})

	// Prefixo de API (exemplo)
	r.Route("/v1/admin", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{"pong": "admin"})
		})
	})

	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	// json.Marshal não deve falhar para mapas simples; ignore erro por simplicidade
	_ = jsonNewEncoder(w).Encode(v)
}

// jsonNewEncoder é isolado para permitir trocar por um encoder custom, se quiser.
func jsonNewEncoder(w http.ResponseWriter) *jsonEncoder {
	return &jsonEncoder{w: w}
}

// Implementação mínima para evitar depender direto de encoding/json aqui em cima.
// (Se preferir, importe "encoding/json" e use direto json.NewEncoder(w).Encode(v))
type jsonEncoder struct{ w http.ResponseWriter }

func (e *jsonEncoder) Encode(v any) error {
	// Import direto aqui para manter o topo do arquivo limpo
	type jsonPkg = struct {
		NewEncoder func(w http.ResponseWriter) *jsonEnc
	}
	// Alias simples de encoding/json
	return jsonNewEncoderStd(e.w).Encode(v)
}

type jsonEnc interface{ Encode(v any) error }

func jsonNewEncoderStd(w http.ResponseWriter) *json.Encoder {
	return json.NewEncoder(w)
}
