FROM alpine:3.22

WORKDIR /app/data

RUN apk add --no-cache ca-certificates tzdata

COPY bin/app /usr/local/bin/app
COPY config/config.example.yaml /app/data/config/config.example.yaml

ENTRYPOINT ["/usr/local/bin/app"]
CMD ["server"]
