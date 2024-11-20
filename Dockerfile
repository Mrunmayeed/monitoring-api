
# Stage 1: Builder Stage
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o monitoring-api .

# Stage 2: Test and Coverage Report Stage
FROM builder AS tester

# Run tests with coverage
RUN mkdir -p /app/coverage && \
    go test ./... -coverprofile=/app/coverage/coverage.out -v && \
    go tool cover -html=/app/coverage/coverage.out -o /app/coverage/coverage.html

# Stage 3: Runtime Stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/monitoring-api .

# Copy the coverage reports (optional, for inspection in artifacts)
COPY --from=tester /app/coverage /app/coverage

RUN mkdir -p temp/log

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./monitoring-api"]
