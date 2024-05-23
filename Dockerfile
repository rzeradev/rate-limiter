FROM golang:1.21.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o server ./cmd/server

# FROM debian:bookworm-slim
# WORKDIR /app
# COPY --from=builder /app/server .
# COPY --from=builder /app/.env .

EXPOSE 8080

ENTRYPOINT ["./server"]