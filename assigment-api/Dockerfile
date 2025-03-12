# Use oficial Golang Image
FROM golang:1.24.0-alpine3.21

# Set working directory
WORKDIR /go/src/app

# Copy all files from current directory to working directory
COPY . .

# Install dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./main"]