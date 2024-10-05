package config

import (
	"flag"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
)

type AgentOptions struct {
	AddressServer  string
	PollInterval   time.Duration
	ReportInterval time.Duration
}

type Config struct {
	Address 				string `env:"ADDRESS"`
	PollInterval 		int64 `env:"POLL_INTERVAL"`
	ReportInterval	int64 `env:"REPORT_INTERVAL"`
}

func Setup() (*AgentOptions, error) {
	var opts AgentOptions
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil{
		return nil, err
	}

	var (
		address string
		pollInterval int64
		reportInterval int64
	)

	flag.StringVar(&address, "a", "localhost:8080", "server address")
	flag.Int64Var(&pollInterval, "p", 2, "frequency of polling metrics")
	flag.Int64Var(&reportInterval, "r", 2, "frequency of sending metrics to the server")
	flag.Parse()

	switch {
	case os.Getenv("ADDRESS") != "":
		opts.AddressServer = cfg.Address
	default:
		opts.AddressServer = address
	}

	switch {
	case os.Getenv("POLL_INTERVAL") != "":
		opts.PollInterval = time.Duration(cfg.PollInterval)
	default:
		opts.PollInterval = time.Duration(pollInterval)
	}

	switch {
	case os.Getenv("REPORT_INTERVAL") != "":
		opts.ReportInterval = time.Duration(cfg.ReportInterval)
	default:
		opts.ReportInterval = time.Duration(reportInterval)
	}

	return &opts, nil
}