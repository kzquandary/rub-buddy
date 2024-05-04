FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

CMD ["./main"]
