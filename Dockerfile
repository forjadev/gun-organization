# Stage 1: Build Go application
FROM golang:1.20-alpine3.17 AS builder

WORKDIR /app
COPY . .

RUN go build -o app

EXPOSE 9000

CMD ["./app"]