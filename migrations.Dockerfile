FROM alpine:3.18.4

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.15.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD .env .
ADD scripts/migration.sh .

RUN chmod +x migration.sh

CMD ["bash", "migration.sh"]