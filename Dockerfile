FROM golang:1.20 AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} cmd/app/main.go

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
EXPOSE 8080
ENV PATH="/go/bin:${PATH}"
CMD ["main"]