package amneziawg

import (
	"fmt"
	"net"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

func (a *awg) AddPeer(fileName string) (string, error) {

	// генерируем виртуальный IP
	peerVirtualEndpoint, err := a.allowedIPS()
	if err != nil {
		return "", err
	}
	// парсим маску и IP клиента
	_, ipNet, err := net.ParseCIDR(peerVirtualEndpoint)
	if err != nil {
		return "", err
	}

	// создание приватного ключа для пользователя
	peerPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", err
	}

	// генерируем PresharedKey
	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return "", err
	}

	// создаем конфигурационный файл для пользователя
	if err := a.createPeerCfg(fileName, peerPrivateKey, presharedKey, peerVirtualEndpoint); err != nil {
		return "", err
	}

	peerPublicKey := peerPrivateKey.PublicKey()

	peerCfg := wgtypes.PeerConfig{
		PublicKey:    peerPublicKey,
		PresharedKey: &presharedKey,
		AllowedIPs:   []net.IPNet{*ipNet},
	}

	cfg := wgtypes.Config{
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{peerCfg},
	}

	// Задаем новую конфигурацию девайса (туннель)
	if err := a.client.ConfigureDevice(a.device.Name, cfg); err != nil {
		return "", err
	}

	return peerPublicKey.String(), nil
}

// Выделение IP с проверкой занятости
// ПОТОМ УДАЛИМ ЭТУ ФУНКЦИЮ
func (a *awg) allowedIPS() (string, error) {

	// Собираем занятые IP
	usedIPs := make(map[string]bool)
	for _, peer := range a.device.Peers {
		for _, allowedIP := range peer.AllowedIPs {
			usedIPs[allowedIP.IP.String()] = true
		}
	}

	// Ищем свободный IP
	for i := 2; i < 255; i++ {
		endpoint := fmt.Sprintf("10.66.66.%d/32", i)
		if !usedIPs[endpoint] {
			return endpoint, nil
		}
	}

	return "", fmt.Errorf("no free IPs available")
}

// ДОП УЛУЧШЕНИЯ
// После успешного добавления сохраните:
// - userPrivateKey.String() (для выдачи клиенту)
// - userPublicKey.String() (для идентификации)
// - usrIPStr (выделенный IP)
// - username
