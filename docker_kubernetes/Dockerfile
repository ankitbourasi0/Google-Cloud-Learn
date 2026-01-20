#reference - https://www.youtube.com/watch?v=S2Atn6LHXeQ

# Use the official Golang image as the base
FROM golang:1.25.5-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the Go binary
RUN go build -o /app .

#-------------------------------------------------------------------------
# Final lightweight stage
FROM alpine:3.21 AS final

# Install ca-certificates for HTTPS connections (Cloud SQL needs this)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app /bin/app

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["/bin/app"]








