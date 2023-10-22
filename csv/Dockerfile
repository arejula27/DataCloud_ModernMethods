# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go application inside the container
RUN go build -o main

# Expose port 1234
EXPOSE 1234

# Run the main command when the container starts
CMD ["./main"]
