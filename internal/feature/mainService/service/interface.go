package service

import (
	"app/internal/feature/mainService/telegram"
)

type awgClient interface {
	AddPeer(fileName, virtualEndpoint string) (string, error)
	DeletePeer(peerPublicKeyStr string) error
}

type service struct {
	tg  *telegram.Tg
	awg awgClient
}

func New(tg *telegram.Tg, awg awgClient) service {

	return service{
		tg:  tg,
		awg: awg,
	}
}
