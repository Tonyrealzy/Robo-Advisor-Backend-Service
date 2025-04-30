# Use a specific version of the Golang base image
FROM golang:1.23.5-alpine

# Set the working directory in the container
WORKDIR /app

# Copy all files from the current directory to the /app directory in the container
COPY . .

# Install dependencies and build the Go application
RUN go mod tidy && go build -o app

# Define the command to run your Go application
CMD ["./app"]
