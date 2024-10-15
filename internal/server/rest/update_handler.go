package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	GaugeTypeStr = string("gauge")
	CounterTypeStr = string("counter")
)

type UpdateService interface {
	AddGauge(string, string) error
	AddCounter(string, string) error
}

type UpdateHandler struct {
	Service UpdateService
}

func NewUpdateHandler(router *chi.Mux, svc UpdateService) {
	handler := &UpdateHandler{
		Service: svc,
	}

	router.Post("/update/{type}/{metric}/{value}", handler.Update)
}

func (s *UpdateHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Проверка метода отправки запроса
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// // Проверка параметра content-type
	// if r.Header.Get("Content-Type") != "text/plain" {
	// 	w.WriteHeader(http.StatusUnsupportedMediaType)
	// 	return
	// }

	metricType := chi.URLParam(r, "type")
	metricName := chi.URLParam(r, "metric")
	metricValue := chi.URLParam(r, "value")

	// Валидация имени метрики
	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Валидация по типу метрики
	switch metricType {
	case GaugeTypeStr:
		err := s.Service.AddGauge(metricName, metricValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		return
	case CounterTypeStr:
		err := s.Service.AddCounter(metricName, metricValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}