package main

import (
	"sync"

	"github.com/sur1k1/go-metrics/internal/agent/metric"
	"github.com/sur1k1/go-metrics/internal/agent/config"
)

func main() {
	flagOpts := config.FlagsOptions()
	// Инициализация временного хранилища метрик
	s := metric.NewMetricStorage()

	var wg sync.WaitGroup
	wg.Add(1)

	go metric.MetricUpdater(s, flagOpts)
	go metric.MetricSender(s, flagOpts)

	wg.Wait()
}