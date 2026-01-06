# ---------- BUILDER -----------
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk update && apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build optimizado y pequeño
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd
# ---------- RUNNER ----------
FROM alpine:latest

WORKDIR /app


RUN apk add --no-cache tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/internal ./internal
ENV TZ=America/Bogota

EXPOSE 3004

CMD ["./main"]

