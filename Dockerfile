FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml .
EXPOSE 8080
CMD ["/app/main"]