# Use the official Golang image as the base image
FROM golang:1.22

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Modules files first
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/gymshark-challenge

# Copy the Vue.js build files into the container
COPY vue-app/dist ./vue-app/dist

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
