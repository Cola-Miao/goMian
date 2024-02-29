FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
COPY . .
RUN go build -o goMian main.go

FROM alpine

WORKDIR /build
COPY --from=builder /build/goMian /build/goMian

CMD ["./goMian"]