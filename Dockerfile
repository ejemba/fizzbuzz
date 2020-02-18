FROM golang:alpine as builder
WORKDIR /fizzbuzz
RUN apk add --no-cache git
COPY . /fizzbuzz
RUN CGO_ENABLED=0 GOOS=linux go build -o fizzbuzz cmd/main.go

FROM alpine:edge
COPY --from=builder /fizzbuzz/fizzbuzz /fizzbuzz
EXPOSE 3000
CMD ["/fizzbuzz"]