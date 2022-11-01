FROM golang:1.18.5-bullseye as builder

WORKDIR /app

COPY . /app

RUN go build -o notification-service . && chmod +x notification-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/notification-service /app

CMD [ "/app/notification-service" ]