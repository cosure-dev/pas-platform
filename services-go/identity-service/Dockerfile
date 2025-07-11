FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy migrations so they are available to the migration tool
COPY migrations ./migrations

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /identity-service ./cmd/main.go

# Create the final, minimal production image
FROM gcr.io/distroless/static-debian11

WORKDIR /

# Copy the migrations from the builder stage
COPY --from=builder /app/migrations ./migrations

# Copy the static binary from the builder stage
COPY --from=builder /identity-service /identity-service

EXPOSE 8081
ENTRYPOINT ["/identity-service"]