package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// описание интерфейса получения метрик из временного
type MetricGetter interface {
	GetMetric(metricType, metricName string) (string, error)
}

// хэндлер обрабатывает входящие данные (тип метрики, имя метрики) и отдает значение
func MetricHandler(m MetricGetter) http.HandlerFunc {
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
		
		value, err := m.GetMetric(metricType, metricName)
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", value)
	}
}