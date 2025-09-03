package main

import (
	"log"
	"net/http"
	"os"

	"palantir/services/api-agent/internal/httpserver"
	"palantir/services/api-agent/internal/store"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	s, err := store.Open(dsn)
	if err != nil { log.Fatalf("db: %v", err) }
	defer s.Close()

	port := os.Getenv("AGENT_API_PORT")
	if port == "" { port = "8082" }
	r := httpserver.Router(s)
	log.Printf("Palantir Agent API :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
