# ----------- BUILDER STAGE -----------
FROM golang:alpine AS builder

RUN apk add --no-cache \
    build-base \
    libwebp-dev 
    
WORKDIR /app

# Go deps
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy source
COPY ./api .

# Generate docs
RUN swag init -g main.go -o ./docs

# Build binary
RUN go build -o main main.go


# ----------- RUNTIME STAGE -----------
FROM alpine:latest

RUN apk add --no-cache \
    libwebp \
    varnish \
    tzdata

WORKDIR /app

# Copy only what we need from builder
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Copy runtime files
COPY default.vcl /etc/varnish/default.vcl
COPY entrypoint.sh ./entrypoint.sh

RUN chmod +x ./entrypoint.sh

CMD ["/bin/sh", "/app/entrypoint.sh"]