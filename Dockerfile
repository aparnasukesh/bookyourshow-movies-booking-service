# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main cmd/main.go

# Stage 2: Create the final image with minimal size
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy environment variables file
COPY .env .

# Expose the port that the app listens on
EXPOSE 5050

# Command to run the application
CMD ["./main"]
