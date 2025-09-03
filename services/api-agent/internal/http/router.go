package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"Palantir/services/api-agent/internal/store"
)

type Event struct {
	TS      string                 `json:"ts"`                 // RFC3339
	Type    string                 `json:"type"`               // ex: "network.connection"
	Source  string                 `json:"source"`             // plugin name
	Severity string                `json:"severity,omitempty"` // info/low/medium/high/critical
	AgentID string                 `json:"agent_id,omitempty"`
	TenantID string                `json:"tenant_id,omitempty"`
	Payload map[string]any         `json:"payload"`
}

func Router(s *store.Store) http.Handler {
	r := chi.NewRouter()

	r.Post("/v1/events", func(w http.ResponseWriter, req *http.Request) {
		var body struct {
			Events []Event `json:"events"`
		}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest); return
		}
		if len(body.Events) == 0 {
			http.Error(w, "empty events", http.StatusBadRequest); return
		}

		batch := make([]store.EventRow, 0, len(body.Events))
		for _, e := range body.Events {
			if e.TS == "" || e.Type == "" || e.Source == "" || e.Payload == nil {
				http.Error(w, "missing required fields", http.StatusBadRequest); return
			}
			ts, err := time.Parse(time.RFC3339, e.TS)
			if err != nil { http.Error(w, "invalid ts", http.StatusBadRequest); return }
			row := store.EventRow{
				TenantID: nullStr(e.TenantID),
				AgentID:  nullStr(e.AgentID),
				Type:     e.Type,
				Source:   e.Source,
				Severity: nullStr(e.Severity),
				TS:       ts,
				Payload:  e.Payload,
			}
			batch = append(batch, row)
		}
		if err := s.InsertEvents(req.Context(), batch); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError); return
		}
		w.WriteHeader(http.StatusAccepted)
	})

	// health
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request){ w.Write([]byte(`{"ok":true}`)) })

	return r
}

func nullStr(s string) *string { if s == "" { return nil }; return &s }
