FROM golang:1.21-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR "/app"

COPY . .

RUN go mod tidy

RUN go build -o binary cmd/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /app .

EXPOSE 7070

ENTRYPOINT ["/binary"]