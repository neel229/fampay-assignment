FROM alpine:latest

COPY fampay .

COPY app.env .

EXPOSE 8080

RUN apk add --no-cache ca-certificates openssl

CMD ["/fampay"]