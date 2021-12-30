FROM alpine:latest

RUN apk add --no-cache tzdata
COPY go-transip-dyndns /usr/bin
CMD ["/usr/bin/go-transip-dyndns", "update", "-k"]