package main

import (
	"os"
	"telegram-service/internal/service"
	"telegram-service/internal/telegram"
)

func main() {
	tgKey := os.Getenv("TELEGRAM_KEY")
	url := os.Getenv("URL")

	if len(tgKey) == 0 && len(url) == 0 {
		panic("TELEGRAM_KEY environment variable is not set")
	}

	tg, err := telegram.New(tgKey)
	if err != nil {
		panic(err)
	}

	client := service.NewHttpClient(url)

	service := service.New(tg, client)

	service.Update()
}
