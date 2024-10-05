package config

import "flag"

func ParseFlags() string {
	var serverAddress string

	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.Parse()

	return serverAddress
}