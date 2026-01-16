package amneziawg

import (
	"fmt"
)

// выводит информацию о работе девайса(интерфейса)
func (a *awg) DeviceInfo() error {

	fmt.Println("----amneziawg работает----")
	fmt.Println("Interface:", a.device.Name)
	fmt.Println("Private key:", a.device.PrivateKey)
	fmt.Println("Public key:", a.device.PublicKey)
	fmt.Println("Listen Port:", a.device.ListenPort)
	fmt.Println("Is amnezia:", a.device.IsAmnezia)

	return nil
}
