# Use the official Go image as the base image for both building and running the app
FROM golang:1.22 AS app

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 (the port your app will listen on)
EXPOSE 8080

# Command to run the Go binary when the container starts
CMD ["./main"]
