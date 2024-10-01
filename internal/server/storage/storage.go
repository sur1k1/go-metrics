package storage

import (
	"strconv"
)

type Storage interface {
	AddGauge(string, string) error
	AddCounter(string, string) error
}

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

	s.GaugeMap[metricName] = floatValue
	return nil
}

// Функция сложения значений COUNTER метрик
func (s *MemStorage) AddCounter(metricName, value string) error {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}

	s.CounterMap[metricName] += intValue
	return nil
}