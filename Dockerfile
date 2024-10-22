# Use an official Go image with Go 1.20
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files for dependency management
COPY go.mod .
COPY go.sum .

# Download Go module dependencies
RUN go mod download

# Copy the rest of your application source code
COPY . .

# Build your Go application
RUN go build -o /dist

# Expose the port your application will listen on
EXPOSE 8001

# Define the command to run your application
CMD [ "/dist" ]
