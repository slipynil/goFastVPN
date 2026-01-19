package main

import (
	"context"
	"os"
	"telegram-service/internal/repository"
	"telegram-service/internal/service"
	"telegram-service/internal/telegram"

	"github.com/jackc/pgx/v5"
)

func main() {
	tgKey := os.Getenv("TELEGRAM_KEY")
	url := os.Getenv("HTTP_URL")
	sql_url := os.Getenv("SQL_URL")

	if len(tgKey) == 0 || len(url) == 0 || len(sql_url) == 0 {
		panic("TELEGRAM_KEY, HTTP_URL, or SQL_URL environment variable is not set")
	}

	// init telegram service
	tg, err := telegram.New(tgKey)
	if err != nil {
		panic(err)
	}

	// init http client service
	client := service.NewHttpClient(url)

	// init postgres service
	conn, err := pgx.Connect(context.Background(), sql_url)
	if err != nil {
		panic(err)
	}
	postgres := repository.New(conn, context.Background())
	if err := postgres.Ping(); err != nil {
		panic(err)
	}

	// init service
	service := service.New(tg, client, postgres)

	// run service
	service.Update()
}
