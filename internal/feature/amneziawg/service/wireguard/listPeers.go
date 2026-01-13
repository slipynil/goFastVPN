package wireguard

import "fmt"

func (s WireGuard) ListPeers() error {
	device, err := s.client.Device(s.deviceName)
	if err != nil {
		return err
	}

	for _, peer := range device.Peers {
		fmt.Println("---PEER CONNECTION---")
		fmt.Println("Virtual Network:", peer.AllowedIPs[0])
		fmt.Println("Public Key:", peer.PublicKey)
		fmt.Println("Last Handshake:", peer.LastHandshakeTime)
	}
	return nil
}
