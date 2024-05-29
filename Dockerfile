# Use an official Golang runtime as a parent image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the local package files to the container’s workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Run the executable
ENTRYPOINT ["./main"]