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
				_, err := s.httpClient.AddPeer("10.66.66.5/32", update.Message.Chat.ID)
				if err != nil {
					fmt.Println("Error adding peer:", err)
				}
				bufer, err := s.httpClient.DownloadConfFile(update.Message.Chat.ID)
				if err != nil {
					fmt.Println(err)
				}
				file := tgbotapi.FileBytes{
					Name:  fmt.Sprintf("%s.conf", update.Message.Chat.UserName),
					Bytes: bufer,
				}
				msg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, file)
				s.tg.Bot.Send(msg)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			s.tg.Bot.Send(msg)

		}
	}
}
