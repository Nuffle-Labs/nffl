FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/plugin/cmd

RUN go build -v -o /usr/local/bin/plugin ./...

FROM debian:bookworm

COPY --from=builder /usr/local/bin/plugin /usr/local/bin/plugin
ENTRYPOINT ["plugin"]
