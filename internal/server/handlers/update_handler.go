package handlers

import (
	"net/http"
	"strings"
)

const (
	Gauge = string("gauge")
	Counter = string("counter")
)

type Storage interface {
	AddGauge(string, string) error
	AddCounter(string, string) error
}

func UpdateHandler(s Storage) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
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
	
		// Парсинг значений в URL
		urlParsed := strings.Split(strings.TrimPrefix(r.URL.String(), "/update/"), "/")
		if len(urlParsed) != 3{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		metricType := urlParsed[0]
		metricName := urlParsed[1]
		metricValue := urlParsed[2]
	
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
}