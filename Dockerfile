# Build stage
FROM golang:1.27.2-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN go build -v -o exemplar ./cmd

# Runtime stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/exemplar .

# Expose port
EXPOSE 3000

# Set environment
ENV PORT=3000

# Run application
CMD ["./exemplar"]
