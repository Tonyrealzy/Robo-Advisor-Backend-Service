# Use a specific version of the Golang base image
FROM golang:1.24.2-alpine

# Set the working directory in the container
WORKDIR /app

# Copy all files from the current directory to the /app directory in the container
COPY . .

# Expose the port the app will run on
EXPOSE 8080

# Install dependencies and build the Go application
RUN go mod tidy && go build -o app

# Define the command to run your Go application
CMD ["./app"]


# # Use a minimal base image with the specific Go version
# FROM golang:1.23.8-alpine

# # Set the working directory inside the container
# WORKDIR /usr/src/app

# # Copy the already built Go binary into the container
# COPY ./app /usr/local/bin/app

# # Expose the port the app will run on
# EXPOSE 8080

# # Set the default command to run the Go binary
# CMD ["/usr/local/bin/app"]
