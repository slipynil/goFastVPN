package wireguard

import (
	"encoding/json"
	"os"

	"github.com/Jipok/wgctrl-go/wgtypes"
)

type obfuscationCfg struct {
	Jc   int
	Jmin int
	Jmax int
	S1   int
	S2   int
	S3   int
	S4   int
}

func (s WireGuard) ConfigureServer() error {

	// === КОНФИГУРАЦИЯ WireGuard КЛИЕНТА ===

	// создание приватного ключа для девайса
	devicePrivKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return err
	}

	port := 5050
	// Junk Packet параметры
	jc := 8
	jmin := 50
	jmax := 1000

	// Message Padding (байты)
	s1 := 32
	s2 := 32
	s3 := 0
	s4 := 0

	obfCfg := obfuscationCfg{jc, jmin, jmax, s1, s2, s3, s4}
	saveCfg(obfCfg)

	// ПРАВИЛО КОНФИГУРАЦИИ
	// nil = поле не изменяется
	// &(zero-value) = очищает значение
	// value = меняет значение

	cfg := wgtypes.Config{
		PrivateKey:   &devicePrivKey,         // приватный ключ девайса
		ListenPort:   &port,                  // порт для подключения к девайсу
		FirewallMark: nil,                    // НЕ ТРОГАТЬ
		ReplacePeers: false,                  // true = удаляет все старые пиры, false просто добавит новые
		Peers:        []wgtypes.PeerConfig{}, // пустые пиры = 0 пользователей
		Jc:           &jc,                    // Count
		Jmin:         &jmin,                  // Min size
		Jmax:         &jmax,                  // Max size

		// Message Padding parameters (bytes)
		S1: &s1, // Init
		S2: &s2, // Response
		S3: &s3, // Cookie
		S4: &s4, // Transport

		// Message Magic Headers
		// In AmneziaWG these can be ranges ("123-456") or single values. Hence string.
		H1: nil, // Init
		H2: nil, // Response
		H3: nil, // Cookie
		H4: nil, // Transport

		// Init Packet Magic / Custom Signature (obfuscation)
		I1: nil,
		I2: nil,
		I3: nil,
		I4: nil,
		I5: nil,
	}

	err = s.client.ConfigureDevice(s.deviceName, cfg)
	return err
}

func saveCfg(cfg obfuscationCfg) error {
	fileName := "obfuscation.txt"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(cfg)
}
