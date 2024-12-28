# Base image with Go >=1.23
FROM golang:1.23

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Copy all files
COPY . .

# Set Air configuration file path
ENV AIR_CONFIG=.air.toml

# Expose the application's default port
EXPOSE 3000

# Run with Air for hot-reloading
CMD ["air"]
