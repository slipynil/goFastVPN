package wireguard

import (
	"app/internal/core/domains"
	"encoding/json"
	"fmt"
	"net"
	"os"
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

	if err := newUserCfg(userPrivateKey, usrIPStr); err != nil {
		return err
	}

	keepAliveEmptyPuckets := time.Second * 25
	peerCfg := wgtypes.PeerConfig{
		PublicKey:                   userPrivateKey.PublicKey(),
		PersistentKeepaliveInterval: &keepAliveEmptyPuckets,
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

func newUserCfg(userPrivateKey wgtypes.Key, userIPstr string) error {
	file, err := os.Open("data/obfuscation.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	cfg := domains.ObfuscationCfg{}
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return err
	}

	fmt.Printf(`
[Interface]
PrivateKey = %s
Address = %s/32
Jc = %v
Jmin = %v
Jmax = %v
S1 = %v
S2 = %v
S3 = %v
S4 = %v
H1 = 1106457265
H2 = 249455488
H3 = 1209847463
H4 = 1646644382

[Peer]
PublicKey = %v
Endpoint = 91.103.140.214/5050
AllowedIPs = 0.0.0.0/0`,
		userPrivateKey,
		userIPstr,
		cfg.Jc,
		cfg.Jmin,
		cfg.Jmax,
		cfg.S1,
		cfg.S2,
		cfg.S3,
		cfg.S4,
		cfg.PublicServerKey,
	)

	return nil
}

// ДОП УЛУЧШЕНИЯ
// После успешного добавления сохраните:
// - userPrivateKey.String() (для выдачи клиенту)
// - userPublicKey.String() (для идентификации)
// - usrIPStr (выделенный IP)
// - username
