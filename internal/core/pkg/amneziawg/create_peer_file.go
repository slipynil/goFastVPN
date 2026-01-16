package amneziawg

import (
	"fmt"
	"os"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

// создает новый конфигурационный файл для подключения пользователя к туннелю
func (a *awg) createPeerCfg(fileName string, peerPrivateKey wgtypes.Key, presharedKey wgtypes.Key, peerVirtualIP string) error {
	publicDeviceKey := a.device.PublicKey.String()

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
		a.obfuscation.Jc,
		a.obfuscation.Jmin,
		a.obfuscation.Jmax,
		a.obfuscation.S1,
		a.obfuscation.S2,
		a.obfuscation.H1,
		a.obfuscation.H2,
		a.obfuscation.H3,
		a.obfuscation.H4,
		publicDeviceKey,
		presharedKey.String(),
		a.endpoint,
	)

	// Создаем файл конфигурации
	// В будущем будем называть файл по ID юзера
	file, err := os.Create("data/" + fileName + ".conf")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(str)); err != nil {
		return err
	}

	return nil
}
