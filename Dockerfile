# Build stage
FROM golang:1.21.2-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build \
    -ldflags="-X 'gitlab.login.no/tekkom/web/beehive/admin-api.version=${IMAGE_VERSION}'" \
    -o bin/main \
    cmd/main.go

# To get the time zone data
FROM alpine:latest as alpine-with-tz
RUN apk --no-cache add tzdata zip
WORKDIR /usr/share/zoneinfo

#Compressing the zone data
RUN zip -q -r -0 /zoneinfo.zip .

# Run stage
FROM scratch

WORKDIR /app

# Setting time zone data
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine-with-tz /zoneinfo.zip /
ENV TZ=Europe/Oslo

# Fetching the cert hints.
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/bin/main .

CMD ["./main"]
