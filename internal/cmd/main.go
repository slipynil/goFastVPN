package main

import (
	getenv "app/internal/core/pkg/getEnv"
	"app/internal/feature/mainService/service"
	"app/internal/feature/mainService/telegram"
	"os"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

func main() {
	// USE values ONLY ACTIVE tunnel's CONFIGURATION
	// also in /etc/amnezia/amneziawg/awg0.conf
	// available only this values
	tg, err := telegram.New(os.Getenv("TELEGRAM_KEY"))
	if err != nil {
		panic(err)
	}

	cfg, err := getenv.NewObfuscation()
	if err != nil {
		panic(err)
	}

	awg, err := awgctrlgo.New(os.Getenv("DEVICE"), os.Getenv("ENDPOINT"), cfg)
	if err != nil {
		panic(err)
	}

	service.New(tg, awg)

	// // client for managing amneziawg devices
	// // Not creating a new tunnel, using existing one
	// defer awg.Close()

	// // information about the tunnel
	// awg.DeviceInfo()

	// // create a new peer
	// userPublicKey, err := awg.AddPeer("user", "10.66.66.02/32")
	// if err != nil {
	// 	panic(err)
	// }
	// // information about the peers
	// awg.ShowPeers()

	// // delete a peer by public key
	// time.Sleep(time.Minute * 5)
	// if err := awg.DeletePeer(userPublicKey); err != nil {
	// 	panic(err)
	// }
}
