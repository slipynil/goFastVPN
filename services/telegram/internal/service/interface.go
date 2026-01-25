package service

import (
	"telegram-service/internal/dto"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type postgres interface {
	Ping() error
	Close() error
	GetHostID(telegramID int64) (int, error)
	AddUser(username string, telegramID int64, expiresAt time.Time) error
	UpdateStatusTrue(telegramID int64) error
	DeleteClient(telegramID int64) error
}

type telegramClient interface {
	// Chan возвращает канал обновлений (от tgbotapi.UpdatesChannel)
	Chan() tgbotapi.UpdatesChannel

	// Menu отправляет сообщение с главным меню
	Menu(chatID int64) error

	// UpdateMainMenu меняет сообщение на главном меню
	UpdateMainMenu(update tgbotapi.Update) error

	// UpdateSendText меняет текст сообщения и ставит меню "назад"
	UpdateSendText(update tgbotapi.Update, text string) error

	// SendFile отправляет файл (конфиг) пользователю
	SendFile(chat *tgbotapi.Chat, buffer []byte) error
}

type httpClient interface {
	AddPeer(hostID int, telegramID int64) (dto.AddPeerResponse, error)
	DeletePeer(publicKey string) error
	DownloadConfFile(telegramID int64) ([]byte, error)
}

type service struct {
	telegram   telegramClient
	httpClient httpClient
	postgres   postgres
}

func New(telegram telegramClient, httpClient httpClient, postgres postgres) service {

	return service{
		telegram:   telegram,
		httpClient: httpClient,
		postgres:   postgres,
	}
}
