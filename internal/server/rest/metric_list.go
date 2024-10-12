package rest

import (
	"fmt"
	"net/http"
)

type AllMetricsGetter interface {
	GetAllMetrics() map[string]string
}

func MetricListHandler(s AllMetricsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseBody := "Metrics List:\n"

		metrics := s.GetAllMetrics()

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