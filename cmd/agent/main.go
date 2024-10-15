package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/sur1k1/go-metrics/internal/agent/config"
	"github.com/sur1k1/go-metrics/internal/agent/metric"
)

func main() {
	gracefulShutdown()
	flagOpts, err := config.Setup()
	if err != nil{
		fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
	}
	fmt.Printf("Agent started with options:\nServer: %s\nPollInterval: %d\nReportInterval: %d", flagOpts.AddressServer, flagOpts.PollInterval, flagOpts.ReportInterval)
	
	// Инициализация временного хранилища метрик
	s := metric.NewMetricStorage()

	var wg sync.WaitGroup
	wg.Add(1)

	go metric.MetricUpdater(s, flagOpts)
	go metric.MetricSender(s, flagOpts)

	wg.Wait()
}

func gracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("Shutting down gracefully")
		os.Exit(0)
	}()
}