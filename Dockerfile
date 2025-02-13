FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Building the binary
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o engine .


FROM alpine:3.20 AS main
COPY --from=builder /app/engine /bin/
WORKDIR /data

LABEL maintainer="Biltu Das <billionto@gmail.com>"
LABEL org.opencontainers.image.version="0.0.1-alpha"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source=https://github.com/BiltuDas1/crawler-engine

CMD [ "engine" ]
