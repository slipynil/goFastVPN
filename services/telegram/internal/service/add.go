package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) add(chat *tgbotapi.Chat) error {
	// add user to database
	err := s.postgres.AddUser(chat.UserName, chat.ID, time.Now())
	if err != nil {
		return err
	}

	// get Host id
	hostID, err := s.postgres.GetHostID(chat.ID)
	if err != nil {
		s.postgres.DeleteClient(chat.ID)
		return fmt.Errorf("error getting host ID: %w", err)
	}
	// change status on True
	if err := s.postgres.UpdateStatusTrue(chat.ID); err != nil {
		s.postgres.DeleteClient(chat.ID)
		return fmt.Errorf("error updating status: %w", err)
	}

	// add peer to host
	_, err = s.httpClient.AddPeer(hostID, chat.ID)
	if err != nil {
		s.postgres.DeleteClient(chat.ID)
		return fmt.Errorf("Error adding peer: %w", err)
	}
	// get http response buffer of config file
	bufer, err := s.httpClient.DownloadConfFile(chat.ID)
	if err != nil {
		s.postgres.DeleteClient(chat.ID)
		return fmt.Errorf("Error downloading config file: %w", err)
	}
	return s.telegram.SendFile(chat, bufer)
}
