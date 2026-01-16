package amneziawg

import "fmt"

func (s *WireGuard) ShowPeers() error {
	if len(s.device.Peers) == 0 {
		fmt.Println("No peers connected")
		return nil
	}

	for _, peer := range s.device.Peers {
		fmt.Println("---PEER CONNECTION---")
		fmt.Println("Public Key:", peer.PublicKey)
		fmt.Println("Last Handshake:", peer.LastHandshakeTime)
	}
	return nil
}
