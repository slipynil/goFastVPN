package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) add(chat *tgbotapi.Chat, price int) error {
	data := time.Now()
	switch price {
	case 20000:
		data = data.Add(30 * time.Hour * 24)
	case 67:
		data = data.Add(time.Hour * 24)
	}

	if err := s.postgres.NewConnection(chat.ID, data); err != nil {
		return fmt.Errorf("fail to add new entity in postgres %w", err)
	}

	hostID, err := s.postgres.GetHostID(chat.ID)
	if err != nil {
		return fmt.Errorf("error getting host ID: %w", err)
	}
	_, err = s.httpClient.AddPeer(hostID, chat.ID)
	if err != nil {
		return fmt.Errorf("error adding peer: %w", err)
	}
	// get http response buffer of config file
	bufer, err := s.httpClient.DownloadConfFile(chat.ID)
	if err != nil {
		return fmt.Errorf("Error downloading config file: %w", err)
	}
	return s.telegram.SendFile(chat, bufer)
}

func (s *service) getConfFile(u tgbotapi.Update) error {
	chat := u.CallbackQuery.Message.Chat
	if !s.postgres.CheckStatus(chat.ID) {
		s.telegram.UpdateSendText(u, "у вас нет подписки")
		return nil
	}
	// get http response buffer of config file
	bufer, err := s.httpClient.DownloadConfFile(chat.ID)
	if err != nil {
		return fmt.Errorf("Error downloading config file: %w", err)
	}
	return s.telegram.SendFile(chat, bufer)
}
