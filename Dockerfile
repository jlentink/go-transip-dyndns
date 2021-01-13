FROM alpine:latest

COPY go-transip-dyndns /usr/bin
RUN echo '*  *  *  *  *    /usr/bin/go-transip-dyndns' > /etc/crontabs/root
CMD crond -l 2 -f