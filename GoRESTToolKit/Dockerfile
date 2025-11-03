# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the full project source code
COPY . .

# Build the Go binary, targeting the cmd/server main package
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myapp ./cmd/server

# Stage 2: Create a lightweight image for the binary
FROM alpine:3.21

# Install ca-certificates for HTTPS support
RUN apk add --no-cache ca-certificates

# Create non-root user for security
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/myapp .

# Change ownership to non-root user
RUN chown appuser:appuser /app/myapp

USER appuser

# Expose port 8080 (default API port)
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
