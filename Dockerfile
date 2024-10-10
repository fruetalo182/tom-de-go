# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Final stage
FROM alpine:latest

# Set a non-root user
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Copy the image file
COPY /media/tom_real.png /app/media/
COPY /media/tom.png /app/media/

# Set environment variables
ENV PORT=8080
ENV IMAGE_PATH=/app/media/tom_real.png

# Use the non-root user
USER appuser

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]