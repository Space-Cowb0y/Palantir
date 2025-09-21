package web

import (
	"embed"
	// "fmt"
	"net/http"

	"github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
	"github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
)

// go:embed ../../webui/*
var webFS embed.FS

type HTTPServer struct {
	cfg *config.Config
	log logging.Logger
	addr string
}

func NewHTTPServer(cfg *config.Config, log logging.Logger) *HTTPServer {
	return &HTTPServer{cfg: cfg, log: log, addr: cfg.HTTP.Listen}
}

func (s *HTTPServer) Addr() string { return s.addr }

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200); _, _ = w.Write([]byte("ok")) })

	// TODO: wire to real registry
	mux.HandleFunc("/api/eyes", func(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "application/json"); w.Write([]byte("[]")) })

	mux.Handle("/", http.FileServer(http.FS(webFS)))
	s.log.Info("HTTP listening", "addr", s.addr)
	return http.ListenAndServe(s.addr, mux)
}