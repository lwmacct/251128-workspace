FROM golang:1.26.1-alpine AS builder

WORKDIR /src

RUN apk add --no-cache ca-certificates tzdata

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app .

FROM alpine:3.22

WORKDIR /app/data

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /out/app /usr/local/bin/app
COPY config/config.example.yaml /app/data/config/config.example.yaml

ENTRYPOINT ["/usr/local/bin/app"]
CMD ["server"]
