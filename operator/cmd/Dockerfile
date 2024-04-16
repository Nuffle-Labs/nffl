FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/operator/cmd

RUN go build -v -o /usr/local/bin/operator ./...

FROM debian:bookworm

COPY --from=builder /usr/local/bin/operator /usr/local/bin/operator
ENTRYPOINT ["operator"]
