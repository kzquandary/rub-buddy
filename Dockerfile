FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

CMD ["./main"]
