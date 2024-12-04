
# Use the official Golang image to build the Go binary
FROM golang:1.20-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
    