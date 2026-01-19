package service

import "telegram-service/internal/telegram"

type service struct {
	tg         *telegram.Tg
	httpClient client
}

func New(tg *telegram.Tg, client client) service {

	return service{
		tg:         tg,
		httpClient: client,
	}
}
