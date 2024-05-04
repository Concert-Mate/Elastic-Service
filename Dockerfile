FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go 

FROM alpine:latest

RUN apk update && \
    apk add --no-cache bash curl nano

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 50051

#CMD ["./main"]
CMD ["tail", "-f", "/dev/null"]
