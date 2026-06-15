FROM golang:1.24-alpine AS builder
LABEL authors="cuel"
WORKDIR /app
COPY . .


RUN go build -o server main.go
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 9091
CMD ["./server"]
