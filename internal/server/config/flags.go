package config

import (
	"flag"
	"os"
)

type ServerConfig struct {
	ServerAddress string
	LogLevel string
}

func ParseFlags() *ServerConfig {
	var serverAddress string
	var logLevel string

	flag.StringVar(&serverAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&logLevel, "l", "info", "log level")
	flag.Parse()

	if addr := os.Getenv("ADDRESS"); addr != ""{
		serverAddress = addr
	}
	if loggerLevel := os.Getenv("LOG_LEVEL"); loggerLevel != ""{
		logLevel = loggerLevel
	}

	return &ServerConfig{
		ServerAddress: serverAddress,
		LogLevel: logLevel,
	}
}