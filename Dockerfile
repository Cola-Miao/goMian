FROM golang:alpine AS builder
LABEL authors="Cola_Miao"

WORKDIR /build

ADD go.mod .
COPY . .
RUN go build -o goMian goMian.go

FROM alpine

WORKDIR /build
COPY --from=builder /build/goMian /build/goMian

CMD ["./goMian"]