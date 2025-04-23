FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/air-verse/air@latest

EXPOSE 8080
EXPOSE 2345
ENTRYPOINT ["/go/bin/air"]
