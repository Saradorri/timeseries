FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

WORKDIR /app/cmd/api

RUN export GOPROXY=https://goproxy.io,direct && CGO_ENABLED=0 go build -o grpc_server .

FROM alpine:latest

WORKDIR /app

ARG APP_GRPC_PORT=5050
ARG APP_METRIC_PORT=2112

EXPOSE ${APP_GRPC_PORT}
EXPOSE ${APP_METRIC_PORT}

COPY --from=builder /app/cmd/api/grpc_server /app/

COPY ./config.yml /app/

CMD ["/app/grpc_server", "-e", "/app/config.yml"]