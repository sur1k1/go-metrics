package main

import (
	"net/http"
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

const (
	Gauge = string("gauge")
	Counter = string("counter")
)

func main() {
	s := NewStorage()

	mux := http.NewServeMux()
	mux.HandleFunc("/update/{type}/{metric}/{value}", s.UpdateHandler)

	if err := run(mux); err != nil{
		panic(err)
	}
}

func run(mux *http.ServeMux) error {
	return http.ListenAndServe(`:8080`, mux)
}

func NewStorage() *MemStorage {
	return &MemStorage{
		GaugeMap: map[string]float64{},
		CounterMap: map[string]int64{},
	}
}

func (s *MemStorage) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода отправки запроса
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Проверка параметра content-type
	if r.Header.Get("Content-Type") != "text/plain" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// Парсинг значение в URL
	metricType := r.PathValue("type")
	metricName := r.PathValue("metric")
	metricValue := r.PathValue("value")

	// Валидация имени метрики
	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Валидация по типу метрики
	switch metricType {
	case Gauge:
		err := s.AddGauge(metricName, metricValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case Counter:
		err := s.AddCounter(metricName, metricValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
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