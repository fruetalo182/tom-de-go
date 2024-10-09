package main

import (
	"log"

	"image-server/internal/config"
	"image-server/internal/server"
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