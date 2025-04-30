# Build stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o robo-advisor-backend-service .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/robo-advisor-backend-service .

EXPOSE 8000

CMD ["./robo-advisor-backend-service"]
