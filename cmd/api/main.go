package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sims-backend/internal/server"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func gracefulShutdown(fiberServer *server.FiberServer, done chan bool) {
	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// Timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exited gracefully")

	// Notify shutdown completion
	done <- true
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system environment variables")
	}

	// Initialize the Fiber server
	server := server.New()

	// Register routes
	server.RegisterFiberRoutes()

	// Create a channel for shutdown signal
	done := make(chan bool, 1)

	// Start the server
	go func() {
		portStr := os.Getenv("PORT")
		if portStr == "" {
			log.Fatal("Environment variable PORT is not set")
		}

		port, err := strconv.Atoi(portStr)
		if err != nil || port < 1 || port > 65535 {
			log.Fatalf("Invalid PORT: %v", err)
		}

		log.Printf("Starting server on port %d", port)
		err = server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("HTTP server error: %s", err)
		}
	}()

	// Handle graceful shutdown
	go gracefulShutdown(server, done)

	// Wait for the shutdown process to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
