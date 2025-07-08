# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the application, creating a static binary.
# CGO_ENABLED=0 is important for creating a static binary without C dependencies.
# -o /policy-service creates the output file in the root directory.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /policy-service .

# Stage 2: Create the final, minimal production image
# Use a "distroless" image for a minimal footprint and improved security.
FROM gcr.io/distroless/static-debian11

# Set the working directory
WORKDIR /

# Copy the static binary from the builder stage
COPY --from=builder /policy-service /policy-service

# Expose the port the application listens on
EXPOSE 8080

# The command to run the application
ENTRYPOINT ["/policy-service"]