package wireguard

import "github.com/Jipok/wgctrl-go/wgtypes"

// удаляет пир из тунеля по публичному ключу пира
func (s WireGuard) DeletePeer(peerPublicKeyStr string) error {
	peerPublicKey, err := wgtypes.ParseKey(peerPublicKeyStr)
	if err != nil {
		return err
	}
	cfg := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{
			// peerConfig - это конфигурация одного пира. В ней указывается, что делать с конкретным пиром
			{
				PublicKey: peerPublicKey, // публичный ключ пира
				Remove:    true,          // true = удалить этот пир
			},
		},
	}
	return s.client.ConfigureDevice(s.deviceName, cfg)
}
