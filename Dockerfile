# Base image
FROM golang:1.21

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/cosmtrek/air@latest

# Copy application files
COPY . .

# Expose the application's default port
EXPOSE 3000

# Run with Air for hot-reloading
CMD ["air"]
