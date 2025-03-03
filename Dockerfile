FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend cmd/main.go
RUN [ -f backend ] || (echo "Error: backend binary not found" && exit 1)

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/backend .

RUN chmod +x /root/backend

EXPOSE 8080

CMD ["./backend", "-storage", "postgres"]