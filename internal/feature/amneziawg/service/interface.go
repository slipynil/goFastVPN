package service

type WireGuard interface {
	ConfigureServer() error
	AddPeer() error
	DeviceInfo() error
	ListPeers() error
}

type Service struct {
	wiregurad WireGuard
}

func NewService(wiregurad WireGuard) Service {
	return Service{wiregurad: wiregurad}
}
