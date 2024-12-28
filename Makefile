# Simple Makefile for a Go project running on Docker

# Build the application
all: docker-build

docker-build:
	@echo "Building the application inside Docker..."
	@docker compose run --rm app go build -o main.exe cmd/api/main.go

# Run the application
docker-run-app:
	@echo "Running the application inside Docker..."
	@docker compose run --rm app go run cmd/api/main.go

# Create all containers (App, DB, Redis, RabbitMQ)
docker-run:
	@echo "Starting all containers (App, DB, Redis, RabbitMQ)..."
	@docker compose up --build

# Shutdown all containers
docker-down:
	@echo "Stopping all containers..."
	@docker compose down

# Test the application
docker-test:
	@echo "Running tests inside Docker..."
	@docker compose run --rm app go test ./... -v

# Integration Tests for the application
docker-itest:
	@echo "Running integration tests inside Docker..."
	@docker compose run --rm app go test ./internal/database -v

# Clean the binary
docker-clean:
	@echo "Cleaning up build artifacts inside Docker..."
	@docker compose run --rm app rm -f main.exe

# Live Reload with Air
docker-watch:
	@echo "Starting live reload with Air inside Docker..."
	@docker compose run --rm app air

.PHONY: all docker-build docker-run-app docker-test docker-clean docker-watch docker-run docker-down docker-itest
