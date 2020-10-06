FROM alpine:latest

COPY ./go-transip-dyndns /usr/local/bin
RUN echo '*  *  *  *  *    /usr/local/bin/go-transip-dyndns' > /etc/crontabs/root
CMD crond -l 2 -f