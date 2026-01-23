FROM golang:1.25.6-bookworm
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o telegram-service main.go
CMD ["./telegram-service"]
