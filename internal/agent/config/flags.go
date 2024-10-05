package config

import "flag"

type AgentOptions struct {
	AddressServer  string
	PollInterval   int64
	ReportInterval int64
}

func FlagsOptions() *AgentOptions {
	var opts AgentOptions
	flag.StringVar(&opts.AddressServer, "a", "localhost:8080", "server address")
	flag.Int64Var(&opts.PollInterval, "p", 2, "frequency of polling metrics")
	flag.Int64Var(&opts.ReportInterval, "r", 10, "frequency of sending metrics to the server")
	flag.Parse()
	return &opts
}