package main

import (
	"log"
	"os"

	"cloud/src/config"
	"cloud/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Initialize Cassandra database connection
	config.InitDB()
	defer config.CloseDB()

	// Set up Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port " + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
