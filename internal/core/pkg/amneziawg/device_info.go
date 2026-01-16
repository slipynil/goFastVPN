package amneziawg

import (
	"fmt"
)

// выводит информацию о работе девайса(интерфейса)
func (s *WireGuard) DeviceInfo() error {

	fmt.Println("----amneziawg работает----")
	fmt.Println("Interface:", s.device.Name)
	fmt.Println("Private key:", s.device.PrivateKey)
	fmt.Println("Public key:", s.device.PublicKey)
	fmt.Println("Listen Port:", s.device.ListenPort)
	fmt.Println("Is amnezia:", s.device.IsAmnezia)

	return nil
}
