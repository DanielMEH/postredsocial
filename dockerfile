# ---------- BUILDER -----------
FROM golang:1.23.6-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build optimizado y pequeño
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# ---------- RUNNER ----------
FROM alpine:latest

WORKDIR /app

# instalar TZ data
RUN apk add --no-cache tzdata

# Copiamos binario y data necesaria
COPY --from=builder /app/main .
COPY --from=builder /app/data ./data
COPY --from=builder /app/public ./public

# Configurar zona horaria (lo que tu proyecto necesita)
ENV TZ=America/Bogota

EXPOSE 3004

CMD ["./cmd/main"]

