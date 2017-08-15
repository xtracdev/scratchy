FROM bash:latest

RUN adduser -s /bin/bash -D user
WORKDIR /home/user
USER user

ADD ca-certificates.crt /etc/ssl/certs/
