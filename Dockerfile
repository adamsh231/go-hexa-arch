# Start from a minimal Golang Alpine image
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main

# Start from scratch
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy the built executable
COPY --from=builder /app/main .

# Command to run the executable
CMD ["/bin/sh"]
