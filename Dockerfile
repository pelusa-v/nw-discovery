# Use the official Golang image as a base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY ./src .

# Build the Go application
RUN go mod init github.com/pelusa-v/nw-discovery
RUN go build ./cmd/main.go

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the executable
# CMD ["./cmd"]