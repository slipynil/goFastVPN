package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *service) Update() error {
	for {
		updates, err := s.tg.GetUpdatesChan(tgbotapi.NewUpdate(0))
		if err != nil {
			return fmt.Errorf("not created updates: %w", err)
		}
		for update := range updates {
			if update.Message == nil {
				continue
			}

			// add peer command send conf file for user
			switch update.Message.Command() {
			case "add":
				msg, err := s.add(update.Message.Chat)
				if err != nil {
					if psErr := s.postgres.DeleteClient(update.Message.Chat.ID); psErr != nil {
						fmt.Println(psErr, err)
					} else {
						fmt.Println(err)
					}
					continue
				}
				s.tg.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
				s.tg.Send(msg)
			}
		}
	}
}
