package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Palantir/services/api-admin/internal/store"
)

type Server struct{ Store *store.Store }

// GET /v1/admin/events?type=&source=&severity=&from=&to=&limit=&offset=
func (s *Server) ListEvents(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	filter := store.EventFilter{
		Type:     q.Get("type"),
		Source:   q.Get("source"),
		Severity: q.Get("severity"),
		From:     parseTime(q.Get("from")),
		To:       parseTime(q.Get("to")),
		Limit:    atoiDefault(q.Get("limit"), 50),
		Offset:   atoiDefault(q.Get("offset"), 0),
	}
	items, total, err := s.Store.QueryEvents(r.Context(), filter)
	if err != nil { http.Error(w, err.Error(), 500); return }
	json.NewEncoder(w).Encode(map[string]any{"items": items, "total": total})
}

func parseTime(s string) *time.Time {
	if s == "" { return nil }
	if t, err := time.Parse(time.RFC3339, s); err == nil { return &t }
	return nil
}
func atoiDefault(s string, d int) int {
	if s == "" { return d }
	if n, err := strconv.Atoi(s); err == nil { return n }
	return d
}

// GET /v1/admin/events/stream (SSE stub)
func (s *Server) StreamEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	fmt.Fprintf(w, "retry: 5000\n\n")
	flusher, _ := w.(http.Flusher)

	for {
		select {
		case <-r.Context().Done():
			return
		case t := <-ticker.C:
			fmt.Fprintf(w, "event: ping\ndata: {\"ts\":\"%s\"}\n\n", t.UTC().Format(time.RFC3339))
			if flusher != nil { flusher.Flush() }
		}
	}
}
