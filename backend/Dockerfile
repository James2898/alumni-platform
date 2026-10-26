# Development Stage with Air live-reload
FROM golang:1.26-alpine AS dev

WORKDIR /app

RUN apk add --no-cache git curl build-base
RUN go install github.com/air-verse/air@v1.61.7

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

# Production Stage
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/api-server ./cmd/api

FROM alpine:3.20 AS prod

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/api-server .

EXPOSE 8080

CMD ["./api-server"]
