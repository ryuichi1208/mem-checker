FROM golang:1.13.4-alpine3.10 AS builder
WORKDIR /app/backend
COPY main.go .
COPY go.mod .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o mem-checker main.go

FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/backend/mem-checker .
