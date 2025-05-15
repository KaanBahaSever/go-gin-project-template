package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"draw/internal/config"
	"draw/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or cannot be loaded: %v", err)
	}

	// Initialize config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup routes
	r := routes.SetupRoutes()

	// Start the server
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server starting on http://localhost%s", serverAddr)
	r.Run(serverAddr)
}
