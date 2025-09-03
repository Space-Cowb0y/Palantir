package main

import (
	"log"
	"net/http"
	"os"
	"Palantir/services/api-admin/internal/http"
	"Palantir/services/api-admin/internal/store"
)

func main(){
	dsn := os.Getenv("DATABASE_URL")
	db, err := store.Open(dsn)
	if err != nil { log.Fatalf("db: %v", err) }
	defer db.Close()

	port := os.Getenv("ADMIN_API_PORT")
	if port == "" { port = "8081" }
	log.Printf("Palantir Admin API :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, httpserver.Router(db.Pool)))
}
