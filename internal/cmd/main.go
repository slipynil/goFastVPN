package main

import (
	"app/internal/core/domains"
	"app/internal/feature/amneziawg/service/wireguard"
	"os"
	"time"

	"github.com/Jipok/wgctrl-go"
)

func main() {
	// клиент для управления WireGuard устройствами
	WgClient, err := wgctrl.New()
	if err != nil {
		panic(err)
	}
	defer WgClient.Close()
	cfg, err := domains.NewObfuscation()
	if err != nil {
		panic(err)
	}

	device := os.Getenv("DEVICE")
	endpoint := os.Getenv("ENDPOINT")

	wgService := wireguard.WireGuardService(device, endpoint, cfg, WgClient)

	// Инфа о туннеле
	if err := wgService.DeviceInfo(); err != nil {
		panic(err)
	}

	// Создание подключения
	userPublicKey, err := wgService.AddPeer()
	if err != nil {
		panic(err)
	}
	// инфа о подключениях
	if err := wgService.ListPeers(); err != nil {
		panic(err)
	}

	// удаление пира по публичному ключу
	time.Sleep(time.Minute * 5)
	wgService.DeletePeer(userPublicKey)
}
