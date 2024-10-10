package main

import (
	"log"

	"tom-de-go/internal/config"
	"tom-de-go/internal/server"
	
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.New(cfg)
	log.Printf("Server starting on http://localhost:%s", cfg.Port)
	log.Fatal(srv.Run())
}