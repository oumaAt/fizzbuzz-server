# ---- Build stage ----
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Cache dependencies separately from source code
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---- Final stage ----
FROM alpine:3.20

RUN adduser -D -u 1000 appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT ["./server"]