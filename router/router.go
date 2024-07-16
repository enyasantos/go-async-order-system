package router

import (
	"log"
	"net/http"
	"os"
)

func Initialize() {
	mx := http.NewServeMux()
	initializeRoutes(mx)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Server listening on %s", addr)

	if err := http.ListenAndServe(addr, mx); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
