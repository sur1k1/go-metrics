package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// описание интерфейса получения метрик из временного
type MetricService interface {
	GetMetric(metricType, metricName string) (string, error)
}

type MetricHandler struct {
	Service MetricService
}

func NewMetricHandler(router *chi.Mux, svc MetricService) {
	handler := &MetricHandler{
		Service: svc,
	}

	router.Get("/value/{type}/{metric}", handler.MetricValue())
}

// хэндлер обрабатывает входящие данные (тип метрики, имя метрики) и отдает значение
func (s *MetricHandler) MetricValue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := chi.URLParam(r, "type")
		if metricType == ""{
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metricName := chi.URLParam(r, "metric")
		if metricName == ""{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		
		value, err := s.Service.GetMetric(metricType, metricName)
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", value)
	}
}