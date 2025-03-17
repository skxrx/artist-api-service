FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]