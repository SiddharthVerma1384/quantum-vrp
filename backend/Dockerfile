FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git build-base

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /app/main ./cmd/api

FROM alpine:3.19

RUN apk add --no-cache ca-certificates curl

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main ./main
COPY --from=builder /src/docs ./docs

# Create non-root user and set permissions
RUN addgroup -S app && adduser -S -G app app && chown app:app /app/main

USER app

ENV ENV=production

EXPOSE 8080

CMD ["/app/main"]