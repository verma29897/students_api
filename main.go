package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/verma29897/students_api/config"
	"github.com/verma29897/students_api/database"
	"github.com/verma29897/students_api/handlers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.MustLoad()

	database.ConnectDB()
	defer database.CloseDB()

	// Initialize the Gin router
	r := gin.Default()

	// Setup routes
	api := r.Group("/api/v1")
	handlers.RegisterStudentRoutes(api)

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		database.CloseDB()
		os.Exit(0)
	}()

	hostname := cfg.HTTPServer.Hostname
	if hostname == "" {
		hostname = "0.0.0.0"
	}

	port := cfg.HTTPServer.Addr
	if port == "" {
		port = "8080"
	}

	address := hostname + ":" + port

	// Start the server
	log.Printf("Starting server on %s...", address)
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//Debug On and off

func setGinMode(env string) {
	switch env {
	case "debug":
		gin.SetMode(gin.DebugMode)
		log.Println("Gin running in Debug mode")
	case "release":
		gin.SetMode(gin.ReleaseMode)
		log.Println("Gin running in Release mode")
	default:
		log.Println("Invalid environment value. Defaulting to Release mode.")
		gin.SetMode(gin.ReleaseMode)
	}
}
