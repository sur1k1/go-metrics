package main

import (
	"sync"

	"github.com/sur1k1/go-metrics/internal/agent/metric"
)

func main() {
	// Инициализация временного хранилища метрик
	s := metric.NewMetricStorage()

	var wg sync.WaitGroup
	wg.Add(1)

	go metric.MetricUpdater(s)
	go metric.MetricSender(s)

	wg.Wait()
}