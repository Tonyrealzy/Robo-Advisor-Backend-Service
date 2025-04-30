# Use a specific version of the Golang base image
FROM golang:1.23.5-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the full project
COPY . .

# Build the app binary
RUN go build -o app .

# Final image
FROM gcr.io/distroless/base-debian10

WORKDIR /app
COPY --from=builder /app/app .

# Expose app port (change if needed)
EXPOSE 8080

# Start the app
CMD ["./app"]
