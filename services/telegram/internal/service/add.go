package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) add(chatID int64, price int) error {
	date := time.Now()
	switch price {
	case 20000:
		date = date.Add(30 * time.Hour * 24)
	case 0:
		date = date.Add(time.Hour * 24)
	}

	if err := s.postgres.NewConnection(chatID, date); err != nil {
		return fmt.Errorf("fail to add new entity in postgres %w", err)
	}

	hostID, err := s.postgres.GetHostID(chatID)
	if err != nil {
		return fmt.Errorf("error getting host ID: %w", err)
	}
	data, err := s.httpClient.AddPeer(hostID, chatID)
	if err != nil {
		return fmt.Errorf("error adding peer: %w, status code: %v, description: %v", err, data.Message.StatusCode, data.Message.Error)
	}
	s.postgres.SaveKey(chatID, data.PublicKey)
	// get http response buffer of config file
	bufer, err := s.httpClient.DownloadConfFile(chatID)
	if err != nil {
		return fmt.Errorf("Error downloading config file: %w", err)
	}
	return s.telegram.SendFile(chatID, bufer)
}

func (s *service) getConfFile(u tgbotapi.Update) error {
	chatID := u.CallbackQuery.Message.Chat.ID
	if !s.postgres.CheckStatus(chatID) {
		s.telegram.UpdateSendText(u, "у вас нет подписки")
		return nil
	}
	// get http response buffer of config file
	bufer, err := s.httpClient.DownloadConfFile(chatID)
	if err != nil {
		return fmt.Errorf("Error downloading config file: %w", err)
	}
	return s.telegram.SendFile(chatID, bufer)
}
