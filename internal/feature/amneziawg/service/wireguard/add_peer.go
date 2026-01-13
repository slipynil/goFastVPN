package wireguard

import (
	"fmt"
	"net"
	"time"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

func (s WireGuard) AddPeer() error {

	// создаем виртуальный IP
	usrIPStr, err := s.allowedIPS()
	if err != nil {
		return err
	}
	// парсим маску и IP клиента
	_, ipNet, err := net.ParseCIDR(usrIPStr + "/32")
	if err != nil {
		return err
	}

	// создание приватного ключа для девайса
	userPrivateKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return err
	}

	keepAliveEmptyPuckets := time.Second * 25
	peerCfg := wgtypes.PeerConfig{
		PublicKey:                   userPrivateKey.PublicKey(),
		Remove:                      false,
		UpdateOnly:                  false,
		PresharedKey:                nil,
		Endpoint:                    nil, // НЕ МЕНЯТЬ
		PersistentKeepaliveInterval: &keepAliveEmptyPuckets,
		ReplaceAllowedIPs:           false,
		AllowedIPs:                  []net.IPNet{*ipNet},
	}

	cfg := wgtypes.Config{
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{peerCfg},
	}

	return s.client.ConfigureDevice(s.deviceName, cfg)
}

// Выделение IP с проверкой занятости
func (s WireGuard) allowedIPS() (string, error) {
	// Получаем текущих пиров
	device, err := s.client.Device(s.deviceName)
	if err != nil {
		return "", err
	}

	// Собираем занятые IP
	usedIPs := make(map[string]bool)
	for _, peer := range device.Peers {
		for _, allowedIP := range peer.AllowedIPs {
			usedIPs[allowedIP.IP.String()] = true
		}
	}

	// Ищем свободный IP
	for i := 2; i < 255; i++ {
		ip := fmt.Sprintf("10.0.0.%d", i)
		if !usedIPs[ip] {
			return ip, nil
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
