FROM golang:1.24-alpine

RUN apk add --no-cache \
    build-base \
    libwebp-dev

WORKDIR /app

COPY ./api/go.mod .
COPY ./api/go.sum .

RUN go mod download

COPY ./api .

RUN go build -o main main.go

ENTRYPOINT [ "./main" ]