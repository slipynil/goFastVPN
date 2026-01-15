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
	wgClient, err := wgctrl.New()
	if err != nil {
		panic(err)
	}
	defer wgClient.Close()
	cfg, err := domains.NewObfuscation()
	if err != nil {
		panic(err)
	}

	deviceName := os.Getenv("DEVICE")
	wgDevice, err := wgClient.Device(deviceName)
	if err != nil {
		panic(err)
	}

	endpoint := os.Getenv("ENDPOINT")

	wgService := wireguard.WireGuardService(endpoint, cfg, wgClient, wgDevice)

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
	if err := wgService.ShowPeers(); err != nil {
		panic(err)
	}

	// удаление пира по публичному ключу
	time.Sleep(time.Minute * 5)
	wgService.DeletePeer(userPublicKey)
}
