package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"palantir/services/api-admin/internal/store"
)

func Router(db *store.Store) http.Handler {
	s := &Server{Store: db}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET","POST","PATCH","DELETE","OPTIONS"},
		AllowedHeaders:   []string{"Authorization","Content-Type"},
		AllowCredentials: true,
	}))

	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request){ w.Write([]byte(`{"ok":true}`)) })

	// MVP: sem JWT aqui; adicione depois o middleware de OIDC/JWKS
	r.Get("/v1/admin/events", s.ListEvents)
	r.Get("/v1/admin/events/stream", s.StreamEvents)

	return r
}
