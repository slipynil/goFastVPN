package amneziawg

import (
	"github.com/Jipok/wgctrl-go"
	"github.com/Jipok/wgctrl-go/wgtypes"
)

// интерфейс для работы с WireGuard
type awgClient interface {
	ConfigureDevice(name string, cfg wgtypes.Config) error
	Device(name string) (*wgtypes.Device, error)
	Close() error
}

type awg struct {
	endpoint    string
	obfuscation Obfuscation
	client      awgClient
	device      *wgtypes.Device
}

// IP:PORT
// config for obfuscation
func New(tunnelName string, endpoint string, obfuscation *Obfuscation) (*awg, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	device, err := client.Device(tunnelName)
	if err != nil {
		return nil, err
	}
	return &awg{
		endpoint:    endpoint,
		obfuscation: *obfuscation,
		client:      client,
		device:      device,
	}, nil
}
