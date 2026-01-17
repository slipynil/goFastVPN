package main

import (
	getenv "app/internal/core/pkg/getEnv"
	"app/internal/feature/mainService/service"
	"app/internal/feature/mainService/telegram"
	"os"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

func main() {
	tgKey := os.Getenv("TELEGRAM_KEY")
	if len(tgKey) == 0 {
		panic("TELEGRAM_KEY environment variable is not set")
	}

	tg, err := telegram.New(tgKey)
	if err != nil {
		panic(err)
	}

	cfg, err := getenv.NewObfuscation()
	if err != nil {
		panic(err)
	}

	awg, err := awgctrlgo.New(os.Getenv("DEVICE"), os.Getenv("ENDPOINT"), cfg)
	if err != nil {
		panic(err)
	}

	service := service.New(tg, awg)

	service.Update()
}
