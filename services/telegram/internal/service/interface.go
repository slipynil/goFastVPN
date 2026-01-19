package service

import "telegram-service/internal/telegram"

type postgres interface {
	Ping() error
}

type service struct {
	tg         *telegram.Tg
	httpClient client
	postgres   postgres
}

func New(tg *telegram.Tg, client client, postgres postgres) service {

	return service{
		tg:         tg,
		httpClient: client,
		postgres:   postgres,
	}
}
