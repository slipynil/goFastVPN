package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	client, err := wgctrl.New()
	if err != nil {
		panic(err)
	}

	// === КОНФИГУРАЦИЯ СЕРВЕРА ===

	// создание приватного ключа сервера
	ServerPrivKey, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		panic(err)
	}
	port := 5050
	firewall := 0

	cfg := wgtypes.Config{
		PrivateKey:   &ServerPrivKey,
		ListenPort:   &port,
		FirewallMark: &firewall,
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{},
	}

	err = client.ConfigureDevice("wg0", cfg)
	if err != nil {
		panic(err)
	}
	device, err := client.Device("wg0")

	fmt.Println("----wireguard работает----")
	fmt.Println("Interface:", "wg0")
	fmt.Println("Private key:", device.PrivateKey)
	fmt.Println("Public key:", device.PublicKey)
	fmt.Println("Listen Port:", device.ListenPort)
}
