# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
# Output directory is ./docs, main.go is in ./cmd/main.go
RUN swag init -g cmd/main.go --output ./docs

# Build the application
# Using -ldflags="-w -s" to reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main cmd/main.go

# Production stage
FROM alpine:latest

# ca-certificates are important for making HTTPS calls from within the container
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Prebuilt binary file from the previous stage
COPY --from=builder /app/main /app/main
# Copy the .env file. Consider managing secrets/configs externally for production.
COPY --from=builder /app/.env /app/.env
# Copy migrations
COPY --from=builder /app/db /app/db
# Copy Swagger docs
COPY --from=builder /app/docs /app/docs

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]