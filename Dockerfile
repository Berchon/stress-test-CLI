# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install CA certificates (needed for HTTPS requests)
RUN apk add --no-cache ca-certificates

# Download dependencies first (cache layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build static binary for Linux amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test ./cmd/cli/main.go

# ---------- Final stage ----------
FROM scratch

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/stress-test .

# Copy CA certificates bundle (needed for HTTPS requests)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the binary
ENTRYPOINT ["./stress-test"]
