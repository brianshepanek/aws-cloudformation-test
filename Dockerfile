# Start from the official Golang base image
FROM golang

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the server will listen on
EXPOSE 80

# Set the command to run the executable when the container starts
CMD ["./main"]
