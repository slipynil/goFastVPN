package amneziawg

import "fmt"

func (a *awg) ShowPeers() error {
	if len(a.device.Peers) == 0 {
		fmt.Println("No peers connected")
		return nil
	}

	for _, peer := range a.device.Peers {
		fmt.Println("---PEER CONNECTION---")
		fmt.Println("Public Key:", peer.PublicKey)
		fmt.Println("Last Handshake:", peer.LastHandshakeTime)
	}
	return nil
}
