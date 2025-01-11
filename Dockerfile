# Use the official Golang image v-1.23 for building the app
FROM golang:1.23.2 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -ldflags="-w -s" -o webapp


# Use a lightweight image to run the app
FROM alpine:3.18.2

# Set metadata
LABEL maintainer="Christos Markos, Konstantinos Petroutsos, Socrates Aggelakopoulos"
LABEL version="1.0"
LABEL description="Dockerized Ascii-Art-Web"

# Install essential certificates
RUN apk add --no-cache ca-certificates

# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

# Copy the compiled Go binary and necessary assets
COPY --from=builder /app/webapp .
COPY --from=builder /app/banners ./banners
COPY --from=builder /app/templates ./templates

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./webapp"]
