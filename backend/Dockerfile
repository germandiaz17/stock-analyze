# Etapa de build
FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main

# Etapa final
FROM gcr.io/distroless/base-debian10

WORKDIR /app
COPY --from=builder /app/main .
COPY mock.json .

ENV PORT=3000

CMD ["./main"]
