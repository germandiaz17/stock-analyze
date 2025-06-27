# Etapa 1: Construcción del binario
FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY mock.json .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Etapa 2: Contenedor final
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
