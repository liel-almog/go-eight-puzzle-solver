# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Set environment variables
ENV GIN_MODE=release

# Copy all files to the container
COPY . .

# Download dependencies and build the Go application
RUN go mod download
RUN go build -o main .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env

# Expose the application port
EXPOSE 5000

# Run the application
CMD [ "./main" ]