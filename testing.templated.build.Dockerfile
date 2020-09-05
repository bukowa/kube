FROM golang@sha256:5219b39d2d6bf723fb0221633a0ff831b0f89a94beb5a8003c7ff18003f48ead as builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go build -mod vendor -o test ./testing/templated/build
RUN ./test
