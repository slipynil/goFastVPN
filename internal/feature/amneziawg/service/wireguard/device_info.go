package wireguard

import (
	"fmt"
)

// выводит информацию о работе девайса(интерфейса)
func (s WireGuard) DeviceInfo() error {

	device, err := s.client.Device(s.deviceName)
	if err != nil {
		return err
	}

	fmt.Println("----wireguard работает----")
	fmt.Println("Interface:", s.deviceName)
	fmt.Println("Private key:", device.PrivateKey)
	fmt.Println("Public key:", device.PublicKey)
	fmt.Println("Listen Port:", device.ListenPort)
	fmt.Println("Is amnezia:", device.IsAmnezia)

	return nil
}
