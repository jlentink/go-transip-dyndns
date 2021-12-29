FROM alpine:latest

COPY go-transip-dyndns /usr/bin
CMD ["/usr/bin/go-transip-dyndns", "update", "-k"]