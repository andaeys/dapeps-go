# Build Stage
FROM golang:1.21.1-alpine AS builder

WORKDIR /app
COPY . .

# Install PostgreSQL client for Alpine Linux
RUN apk update && \
    apk add postgresql-client

# Copy the entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN go build -o main .

# Final Stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080

CMD ["/entrypoint.sh"]
