# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building the application..."
	@go build -o main.exe cmd/main.go

# Run the application
run:
	@echo "Running the application..."
	@go run cmd/main.go

# Create all containers (App, DB, Redis, RabbitMQ)
docker-run:
	@echo "Starting all containers (App, DB, Redis, RabbitMQ)..."
	@docker compose up --build

# Shutdown all containers
docker-down:
	@echo "Stopping all containers..."
	@docker compose down

# Test the application
test:
	@echo "Running tests..."
	@go test ./... -v

# Integration Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning up build artifacts..."
	@rm -f main.exe

# Live Reload with Air
watch:
	@echo "Starting live reload with Air..."
	@if [ -x "$(command -v air)" ]; then \
		air; \
	else \
		echo "Air not found, installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

.PHONY: all build run test clean watch docker-run docker-down itest
