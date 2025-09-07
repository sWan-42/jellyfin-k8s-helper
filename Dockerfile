FROM golang:1.25-alpine
LABEL org.opencontainers.image.source https://github.com/sWan-42/jellyfin-k8s-helper
WORKDIR /app
COPY cmd/ .

RUN go build -o sidecar main.go

CMD ["./sidecar"]
