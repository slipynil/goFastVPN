package main

import (
	"app/internal/feature/amneziawg/service/wireguard"

	"golang.zx2c4.com/wireguard/wgctrl"
)

func main() {
	// клиент для управления WireGuard устройствами
	WgClient, err := wgctrl.New()
	if err != nil {
		panic(err)
	}
	defer WgClient.Close()

	wgService := wireguard.WireGuardService("wg0", "127.0.0.1:5050", "5050", WgClient)
	if err := wgService.ConfigureServer(); err != nil {
		panic(err)
	}
	if err := wgService.DeviceInfo(); err != nil {
		panic(err)
	}
	for range 10 {
		if err := wgService.AddPeer(); err != nil {
			panic(err)
		}
	}
	if err := wgService.ListPeers(); err != nil {
		panic(err)
	}
}
