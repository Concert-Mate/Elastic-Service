FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go 

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 50051

CMD ["./main"]
