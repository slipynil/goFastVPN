package main

import (
	"app/internal/feature/amneziawg/service/wireguard"
	"fmt"
	"time"

	"github.com/Jipok/wgctrl-go"
)

func main() {
	time.Sleep(time.Second * 10)
	// клиент для управления WireGuard устройствами
	WgClient, err := wgctrl.New()
	if err != nil {
		panic(err)
	}
	defer WgClient.Close()

	defer func() {
		if r := recover(); r != nil {
			time.Sleep(100 * time.Second)
			fmt.Println("ошибка", r)
		}
	}()

	wgService := wireguard.WireGuardService("awg0", "127.0.0.1:5050", "5050", WgClient)
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
