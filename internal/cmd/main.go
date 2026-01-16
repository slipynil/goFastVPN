package main

import (
	getenv "app/internal/core/pkg/getEnv"
	"os"
	"time"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

func main() {
	// USE values ONLY ACTIVE tunnel's CONFIGURATION
	// also in /etc/amnezia/amneziawg/awg0.conf
	// available only this values

	cfg, err := getenv.NewObfuscation()
	if err != nil {
		panic(err)
	}
	tunnelName := os.Getenv("DEVICE")
	endpoint := os.Getenv("ENDPOINT")

	// client for managing amneziawg devices
	// Not creating a new tunnel, using existing one
	awg, err := awgctrlgo.New(tunnelName, endpoint, cfg)
	if err != nil {
		panic(err)
	}
	defer awg.Close()

	// information about the tunnel
	awg.DeviceInfo()

	// create a new peer
	userPublicKey, err := awg.AddPeer("user", "10.66.66.02/32")
	if err != nil {
		panic(err)
	}
	// information about the peers
	awg.ShowPeers()

	// delete a peer by public key
	time.Sleep(time.Minute * 5)
	if err := awg.DeletePeer(userPublicKey); err != nil {
		panic(err)
	}
}
