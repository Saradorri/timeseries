FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN export GOPROXY=https://goproxy.io,direct && CGO_ENABLED=0 go build -o bin/grpc-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /grpc-server .

ARG APP_GRPC_PORT=5050
ARG APP_METRIC_PORT=2112

EXPOSE ${APP_GRPC_PORT}
EXPOSE ${APP_METRIC_PORT}

CMD ["/app/grpc-server", "-e", "/app/config.yml"]