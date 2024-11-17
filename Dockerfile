FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY cmd ./cmd/
COPY internal/ ./internal/

RUN go build -o main ./cmd/main/main.go

CMD [ "main.exe" ]