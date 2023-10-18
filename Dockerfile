FROM golang:1.21.3-alpine3.18 AS builder
COPY . /github.com/sornick01/auth/source
WORKDIR /github.com/sornick01/auth/source
RUN go build -o ./bin/auth cmd/main.go

FROM alpine:3.18.4
WORKDIR /root/
COPY --from=builder /github.com/sornick01/auth/source/bin/auth .
CMD ["./auth"]
