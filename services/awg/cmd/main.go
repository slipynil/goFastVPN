package main

import (
	getenv "awg-service/internal/getEnv"
	"awg-service/internal/transport"
	"os"
	"path/filepath"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

var (
	DefaultHTTP   = ""
	DefaultDEVICE = ""
	DefaultAWG    = ""
)

func main() {
	cfg, err := getenv.NewObfuscation()

	if err != nil {
		panic(err)
	}

	tunnelName, awgEndpoint := getOpt(os.Getenv("DEVICE"), DefaultDEVICE), getOpt(os.Getenv("AWG_ENDPOINT"), DefaultAWG)
	httpEndpoint := getOpt(os.Getenv("HTTP_ENDPOINT"), DefaultHTTP)

	if tunnelName == "" || awgEndpoint == "" || httpEndpoint == "" {
		panic("DEVICE and AWG_ENDPOINT environment variables are required")
	}

	storagePath, err := filepath.Abs("/etc/amnezia/amneziawg/configs/")

	if err != nil {
		panic(err)
	}

	awg, err := awgctrlgo.New(tunnelName, awgEndpoint, storagePath, cfg)

	if err != nil {
		panic(err)
	}
	service := transport.New(awg, storagePath)
	service.Start(httpEndpoint)
}

func getOpt(value string, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
