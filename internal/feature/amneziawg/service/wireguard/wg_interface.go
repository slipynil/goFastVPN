package wireguard

import "github.com/Jipok/wgctrl-go/wgtypes"

// интерфейс для работы с WireGuard
type WireGuardClient interface {
	Close() error
	ConfigureDevice(name string, cfg wgtypes.Config) error
	Device(name string) (*wgtypes.Device, error)
	Devices() ([]*wgtypes.Device, error)
}

type WireGuard struct {
	deviceName string
	endpoint   string
	port       string
	client     WireGuardClient
}

func WireGuardService(deviceName, endpoint, port string, client WireGuardClient) WireGuard {
	return WireGuard{
		deviceName: deviceName,
		endpoint:   endpoint,
		port:       port,
		client:     client,
	}
}
