package storage

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Gauge = string("gauge")
	Counter = string("counter")
)

type MemStorage struct {
	GaugeMap   map[string]float64
	CounterMap map[string]int64
}

// Получение MemStorage клиента
func NewStorage() *MemStorage {
	return &MemStorage{
		GaugeMap: map[string]float64{},
		CounterMap: map[string]int64{},
	}
}

// Функция обновления значений GAUGE метрик
func (s *MemStorage) AddGauge(metricName, value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	s.GaugeMap[strings.ToLower(metricName)] = floatValue
	return nil
}

// Функция сложения значений COUNTER метрик
func (s *MemStorage) AddCounter(metricName, value string) error {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}

	s.CounterMap[strings.ToLower(metricName)] += intValue
	return nil
}

func (s *MemStorage) GetMetric(metricType, metricName string) (string, error) {
	switch metricType{
	case Gauge:
		value, ok := s.GaugeMap[strings.ToLower(metricName)]
		if !ok {
			return "", errors.New("metric not found")
		}

		return strconv.FormatFloat(value, 'f', -1, 64), nil
	case Counter:
		value, ok := s.CounterMap[strings.ToLower(metricName)]
		if !ok {
			return "", errors.New("metric not found")
		}

		return strconv.FormatInt(value, 10), nil
	}

	return "", errors.New("invalid metric type")
}

func (s *MemStorage) GetAllMetrics() map[string]string {
	metrics := make(map[string]string)

	for name, value := range s.GaugeMap{
		metrics[name] = strconv.FormatFloat(value, 'f', -1, 64)
	}

	for name, value := range s.CounterMap{
		metrics[name] = strconv.FormatInt(value, 10)
	}

	return metrics
}