FROM golang:1.24-alpine

RUN apk add --no-cache \
    build-base \
    libwebp-dev \
    varnish

WORKDIR /app

COPY default.vcl /etc/varnish/default.vcl

COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY entrypoint.sh ./entrypoint.sh

RUN go mod download

COPY ./api .

RUN swag init -g main.go -o ./docs

RUN go build -o main main.go

CMD [ "/bin/sh", "/app/entrypoint.sh" ]