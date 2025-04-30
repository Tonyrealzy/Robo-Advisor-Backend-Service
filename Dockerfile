FROM golang:1.23.5-alpine AS builder

# Enable Go modules
ENV GO111MODULE=on

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the app binary
RUN go build -o robo-advisor-backend-service ./main.go

# Stage 2: Run the built binary in a lightweight image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/robo-advisor-backend-service .

# Expose the port (Render expects your app to listen on PORT 10000)
EXPOSE 8000

# Run the binary
CMD ["./robo-advisor-backend-service"]
