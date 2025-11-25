FROM golang:1.24-alpine

RUN apk add --no-cache \
    build-base \
    libwebp-dev \
    varnish

WORKDIR /app

COPY default.vcl /etc/varnish/default.vcl

COPY ./api/go.mod .
COPY ./api/go.sum .

COPY entrypoint.sh ./entrypoint.sh

RUN go mod download

COPY ./api .

RUN go build -o main main.go

CMD [ "/bin/sh", "/app/entrypoint.sh" ]