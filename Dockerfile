# Use an official Go runtime as a parent image
FROM golang:1.17-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the Go modules file
COPY go.mod .

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o badsec .

# Use the minimal Alpine Linux image as the base image for the final image
FROM alpine:3.14

# Copy the application binary from the builder image
COPY --from=builder /app/badsec /app/badsec

COPY templates/ /templates
COPY static/ /static

# Set the command to run when the container starts
CMD ["/app/badsec"]
