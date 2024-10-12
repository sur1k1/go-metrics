package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MetricListService interface {
	GetAllMetrics() map[string]string
}

type MetricListHandler struct {
	Service MetricListService
}

func NewMetricListHandler(router *chi.Mux, svc MetricListService) {
	handler := &MetricListHandler{
		Service: svc,
	}

	router.Get("/", handler.ListMetrics())
}

func (s *MetricListHandler) ListMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseBody := "Metrics List:\n"

		metrics := s.Service.GetAllMetrics()

		for name, value := range metrics {
			responseBody += fmt.Sprintf("%s: %s\n", name, value)
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(responseBody))
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}