package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *service) Update() error {
	for {
		updates, err := s.tg.Bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
		if err != nil {
			return fmt.Errorf("not created updates: %w", err)
		}
		for update := range updates {
			if update.Message == nil {
				continue
			}

			// add peer command send conf file for user
			if update.Message.IsCommand() && update.Message.Command() == "add" {
				data, err := s.httpClient.AddPeer("10.66.66.5/32", update.Message.Chat.UserName)
				if err != nil {
					fmt.Println("Error adding peer:", err)
				}
				msg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, data.FilePath)
				s.tg.Bot.Send(msg)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			s.tg.Bot.Send(msg)

		}
	}
}
