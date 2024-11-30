# Use the official Go image for building the application
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Use a minimal image to run the application
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the default port
EXPOSE 7777

# Command to run the application
CMD ["./app"]
