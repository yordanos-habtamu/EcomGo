# Stage 1: Build the binary
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Copy migration files into the image (assuming migrations are in the `migrations` directory)
COPY cmd/migrate/migrations /app/migrations/

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/main.go

# Stage 2: Use a minimal image for runtime
FROM alpine:latest

WORKDIR /root/

# Install required certificates and libraries for Alpine
RUN apk add --no-cache ca-certificates

# Install `bash` and `migrate` (for running the migrations)
RUN apk add --no-cache bash && \
    apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

# Copy the binary and migration files from the builder stage
COPY --from=builder /app/myapp .
COPY --from=builder /app/migrations /app/migrations


# Copy the entrypoint script
COPY entrypoint.sh .

# Ensure the binary and script are executable
RUN chmod +x myapp entrypoint.sh

# Expose the application's port
EXPOSE 8080

# Use the entrypoint script
CMD ["sh", "entrypoint.sh"]

