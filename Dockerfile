# Use the official Go 1.22.0 image as a base
FROM golang:1.22.0-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./myapp"]
