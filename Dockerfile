# builder image
FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o run_server ./cmd/server

# run image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/run_server .

EXPOSE 8080

CMD ["./run_server"]