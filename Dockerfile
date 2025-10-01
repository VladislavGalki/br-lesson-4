FROM golang:1.25-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o taskmanager cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/taskmanager .
EXPOSE 8080
CMD ["./taskmanager"]