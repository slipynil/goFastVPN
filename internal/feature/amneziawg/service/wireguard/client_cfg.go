package wireguard

import "github.com/Jipok/wgctrl-go/wgtypes"

func (s WireGuard) ConfigureServer() error {

	// === КОНФИГУРАЦИЯ WireGuard КЛИЕНТА ===

	// создание приватного ключа для девайса
	devicePrivKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return err
	}
	port := 5050

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
	}

	err = s.client.ConfigureDevice(s.deviceName, cfg)
	return err
}
