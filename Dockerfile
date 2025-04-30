# Use a minimal base image with the specific Go version
FROM golang:1.2.35-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the already built Go binary into the container
COPY ./app /usr/local/bin/app

# Expose the port the app will run on
EXPOSE 8080

# Set the default command to run the Go binary
CMD ["/usr/local/bin/app"]
