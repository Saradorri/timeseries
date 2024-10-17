FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

WORKDIR /app/cmd/api

RUN export GOPROXY=https://goproxy.io,direct && CGO_ENABLED=0 go build -o grpc_server .

FROM alpine:latest

WORKDIR /app

EXPOSE 5050

COPY --from=builder /app/cmd/api/grpc_server /app/

COPY ./config.yml /app/

CMD ["/app/grpc_server", "-e", "/app/config.yml"]