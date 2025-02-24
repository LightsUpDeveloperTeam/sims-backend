# Build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy the entire project and build the binary
COPY . .
RUN CGO_ENABLED=0 go build -o /app/bin/api ./cmd/api/main.go

# Final stage
FROM alpine:latest
WORKDIR /root/
# Copy the compiled binary from the builder stage
COPY --from=builder /app/bin/api .
EXPOSE 3000
CMD ["./api"]
