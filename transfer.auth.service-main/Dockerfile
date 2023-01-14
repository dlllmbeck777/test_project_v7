FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o transferAuth ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/transferAuth .
COPY --from=builder /app/config /app/config
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
CMD ["/app/transferAuth","-m=dev"]