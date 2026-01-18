package main

import (
	getenv "awg-service/internal/getEnv"
	"awg-service/internal/transport"
	"os"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

func main() {
	cfg, err := getenv.NewObfuscation()
	if err != nil {
		panic(err)
	}
	tunnelName, endpoint := os.Getenv("DEVICE"), os.Getenv("AWG_ENDPOINT")
	if tunnelName == "" || endpoint == "" {
		panic("DEVICE and AWG_ENDPOINT environment variables are required")
	}
	awg, err := awgctrlgo.New(tunnelName, endpoint, cfg)
	if err != nil {
		panic(err)
	}
	service := transport.New(awg)
	service.Start("HTTP_ENDPOINT")
}
