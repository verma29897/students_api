package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/verma29897/students-api/config"
	"github.com/verma29897/students-api/database"
	"github.com/verma29897/students-api/handlers"
)

func main() {
	// Load the configuration
	cfg := config.MustLoad()

	// Connect to the database
	database.ConnectDB(cfg.Database)

	// Defer closing the database connection
	defer database.CloseDB()

	// Initialize router
	r := gin.Default()

	// Setup routes
	api := r.Group("/api/v1")
	handlers.RegisterStudentRoutes(api)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		database.CloseDB()
		os.Exit(0)
	}()

	// Start the server
	port := cfg.HTTPServer.Addr
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
