package service

import (
	"telegram-service/internal/dto"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type postgres interface {
	Ping() error
	Close() error
	GetHostID(telegramID int64) (int, error)
	AddUser(username string, id int64, expiresAt time.Time) error
	UpdateStatusTrue(telegramID int64) error
	DeleteClient(telegramID int64) error
}

type httpClient interface {
	AddPeer(hostID int, telegramID int64) (dto.AddPeerResponse, error)
	DeletePeer(publicKey string) error
	DownloadConfFile(telegramID int64) ([]byte, error)
}

type service struct {
	tg         *tgbotapi.BotAPI
	httpClient httpClient
	postgres   postgres
}

func New(tg *tgbotapi.BotAPI, httpClient httpClient, postgres postgres) service {

	return service{
		tg:         tg,
		httpClient: httpClient,
		postgres:   postgres,
	}
}
