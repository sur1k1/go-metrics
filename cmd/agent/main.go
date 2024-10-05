package main

import (
	"fmt"
	"sync"

	"github.com/sur1k1/go-metrics/internal/agent/config"
	"github.com/sur1k1/go-metrics/internal/agent/metric"
)

func main() {
	flagOpts, err := config.Setup()
	if err != nil{
		panic(err)
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