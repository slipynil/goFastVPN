package amneziawg

import "github.com/Jipok/wgctrl-go/wgtypes"

// удаляет пир из тунеля по публичному ключу пира
func (a *awg) DeletePeer(peerPublicKeyStr string) error {
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
	return a.client.ConfigureDevice(a.device.Name, cfg)
}
