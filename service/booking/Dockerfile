FROM golang:1.24.3-alpine@sha256:ef18ee7117463ac1055f5a370ed18b8750f01589f13ea0b48642f5792b234044 AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache git

COPY service/booking/go.mod service/booking/go.sum ./
COPY shared ../shared

RUN go mod download

COPY service/booking/. .

RUN go build -o evently-event ./cmd/main.go

FROM alpine:latest@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463e380b75e2e74aff4511df3ef88c

WORKDIR /root/

COPY --from=builder /app/evently-event .

COPY ./service/booking/.env .

EXPOSE 8081

CMD ["./evently-event"]