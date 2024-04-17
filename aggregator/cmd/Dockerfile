FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/aggregator/cmd

RUN go build -v -o /usr/local/bin/aggregator ./...

FROM debian:bookworm

RUN apt update && apt install -yy openssl ca-certificates
RUN update-ca-certificates

COPY --from=builder /usr/local/bin/aggregator /usr/local/bin/aggregator
ENTRYPOINT ["aggregator"]
