FROM alpine:latest

ADD /certs/tls-ca-bundle.pem /etc/ssl/certs/ca-certificates.crt
ADD main /

EXPOSE 3000

CMD ["/main"]