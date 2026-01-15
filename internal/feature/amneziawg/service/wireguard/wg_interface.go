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
	endpoint    string
	obfuscation *domains.Obfuscation
	client      WireGuardClient
	device      *wgtypes.Device
}

func WireGuardService(
	endpoint string, // IP:PORT
	obfuscation *domains.Obfuscation, // config for obfuscation
	client WireGuardClient, // client for WireGuard
	device *wgtypes.Device, // device for WireGuard
) *WireGuard {
	return &WireGuard{
		endpoint:    endpoint,
		obfuscation: obfuscation,
		client:      client,
		device:      device,
	}
}
