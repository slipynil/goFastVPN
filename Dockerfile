FROM golang:1.25.5-bookworm
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/bin ./internal/cmd/main.go
CMD ["./bin"]
