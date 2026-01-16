package main

import (
	"app/internal/core/pkg/amneziawg"
	"os"
	"time"
)

func main() {
	// клиент для управления WireGuard устройствами
	cfg, err := amneziawg.NewObfuscation()
	if err != nil {
		panic(err)
	}

	tunnelName := os.Getenv("DEVICE")
	endpoint := os.Getenv("ENDPOINT")

	amnezia, err := amneziawg.New(tunnelName, endpoint, cfg)
	if err != nil {
		panic(err)
	}
	defer amnezia.Close()

	// Инфа о туннеле
	if err := amnezia.DeviceInfo(); err != nil {
		panic(err)
	}

	// Создание подключения
	userPublicKey, err := amnezia.AddPeer("user")
	if err != nil {
		panic(err)
	}
	// инфа о подключениях
	if err := amnezia.ShowPeers(); err != nil {
		panic(err)
	}

	// удаление пира по публичному ключу
	time.Sleep(time.Minute * 5)
	amnezia.DeletePeer(userPublicKey)
}
