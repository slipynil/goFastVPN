package wireguard

import (
	"app/internal/core/domains"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// интерфейс для работы с WireGuard
type WireGuardClient interface {
	ConfigureDevice(name string, cfg wgtypes.Config) error
	Device(name string) (*wgtypes.Device, error)
}

type WireGuard struct {
	deviceName  string
	endpoint    string
	obfuscation *domains.Obfuscation
	client      WireGuardClient
}

func WireGuardService(deviceName, endpoint string, obfuscation *domains.Obfuscation, client WireGuardClient) *WireGuard {
	return &WireGuard{
		deviceName:  deviceName,
		endpoint:    endpoint,
		obfuscation: obfuscation,
		client:      client,
	}
}
