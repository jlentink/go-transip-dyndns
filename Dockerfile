FROM alpine:latest

COPY go-transip-dyndns /usr/bin
RUN echo '*  *  *  *  * /usr/bin/go-transip-dyndns && echo .' > /etc/crontabs/root
CMD ["crond", "-f", "-l", "2"]