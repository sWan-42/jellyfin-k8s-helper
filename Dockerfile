FROM golang:1.25-alpine

WORKDIR /app
COPY cmd/ .

RUN go build -o sidecar main.go

CMD ["./sidecar"]
