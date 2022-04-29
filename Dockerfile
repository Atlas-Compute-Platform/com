# Atlas Dockerfile
FROM alpine:latest
RUN apk add go
COPY ./com /usr/local/bin/com
EXPOSE 8800/tcp
CMD ["/usr/local/bin/com"]
