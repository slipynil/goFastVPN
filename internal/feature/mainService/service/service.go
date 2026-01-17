package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *service) Update() error {
	updates, err := s.tg.Configure()
	if err != nil {
		return fmt.Errorf("not created updates: %w", err)
	}
	for _, update := range updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			s.tg.Bot.Send(msg)

		}
	}
	return nil
}
