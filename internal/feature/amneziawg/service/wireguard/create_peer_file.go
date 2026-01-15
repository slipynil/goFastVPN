package wireguard

import (
	"fmt"
	"os"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// создает новый конфигурационный файл для подключения пользователя к туннелю
func (s WireGuard) createPeerCfg(peerPrivateKey wgtypes.Key, presharedKey wgtypes.Key, peerVirtualIP string) error {
	device, err := s.client.Device(s.deviceName)
	if err != nil {
		return err
	}
	publicDeviceKey := device.PublicKey.String()

	str := fmt.Sprintf(`
[Interface]
PrivateKey = %s
Address = %s
Jc = %v
Jmin = %v
Jmax = %v
S1 = %v
S2 = %v
H1 = %v
H2 = %v
H3 = %v
H4 = %v

[Peer]
PublicKey = %v
PresharedKey = %v
Endpoint = %v
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`,
		peerPrivateKey,
		peerVirtualIP,
		s.obfuscation.Jc,
		s.obfuscation.Jmin,
		s.obfuscation.Jmax,
		s.obfuscation.S1,
		s.obfuscation.S2,
		s.obfuscation.H1,
		s.obfuscation.H2,
		s.obfuscation.H3,
		s.obfuscation.H4,
		publicDeviceKey,
		presharedKey.String(),
		s.endpoint,
	)

	// Создаем файл конфигурации
	// В будущем будем называть файл по ID юзера
	file, err := os.Create("data/user.conf")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(str)); err != nil {
		return err
	}

	return nil
}
