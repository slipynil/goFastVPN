package wireguard

import "fmt"

func (s WireGuard) ShowPeers() error {
	device, err := s.client.Device(s.deviceName)
	if err != nil {
		return err
	}
	if len(device.Peers) == 0 {
		fmt.Println("No peers connected")
		return nil
	}

	for _, peer := range device.Peers {
		fmt.Println("---PEER CONNECTION---")
		fmt.Println("Public Key:", peer.PublicKey)
		fmt.Println("Last Handshake:", peer.LastHandshakeTime)
	}
	return nil
}
