FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/relayer/cmd

RUN go build -v -o /usr/local/bin/relayer ./...

FROM debian:bookworm

RUN apt-get update && apt-get install -y npm
RUN npm install -g near-cli@3.0.0

COPY --from=build /usr/local/bin/relayer /usr/local/bin/relayer
ENTRYPOINT [ "relayer"]