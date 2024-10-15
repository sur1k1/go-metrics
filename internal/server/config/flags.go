package config

import (
	"flag"
	"os"
)

func ParseFlags() string {
	var serverAddress string

	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.Parse()

	if addr := os.Getenv("ADDRESS"); addr != ""{
		serverAddress = addr
	}

	return serverAddress
}