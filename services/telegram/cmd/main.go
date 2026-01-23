package main

import (
	"context"
	"os"
	httpclient "telegram-service/internal/httpClient"
	"telegram-service/internal/repository"
	"telegram-service/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	tgKey := os.Getenv("TELEGRAM_KEY")
	url := os.Getenv("HTTP_URL")
	dbConn := os.Getenv("DB_CONN")

	if len(tgKey) == 0 || len(url) == 0 || len(dbConn) == 0 {
		panic("TELEGRAM_KEY, HTTP_URL, or DB_CONN environment variable is not set")
	}

	// init telegram service

	tg, err := tgbotapi.NewBotAPI(tgKey)
	if err != nil {
		panic(err)
	}

	// init http client service
	client := httpclient.New(url)

	// init postgres service
	postgres, err := repository.New(context.Background(), dbConn)
	if err != nil {
		panic(err)
	}

	// init service
	service := service.New(tg, client, postgres)

	// run service
	service.Update()
}
