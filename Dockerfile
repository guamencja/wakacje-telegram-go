# Build stage
FROM golang:1.21-alpine AS builder

# Install git (required for go mod)
RUN apk add --no-cache git

WORKDIR /app

# Copy and download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o bot .

# Runtime stage
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/bot .

# Command to run with auto-restart
CMD while true; do ./bot; sleep 5; done
