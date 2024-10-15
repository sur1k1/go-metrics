package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/sur1k1/go-metrics/internal/server/repository/memstorage"
)

func TestMetricListHandler(t *testing.T) {
	tests := []struct {
		name       string
		args       MetricListService
		method		string
		url				string
		wantStatus int
	}{
		{
			name: "positive test #1",
			args: &storage.MemStorage{
				GaugeMap: map[string]float64{
					"alloc": 123.0,
				},
				CounterMap: map[string]int64{
					"pollcount": 123,
				},
			},
			method: "GET",
			url: "/",
			wantStatus: 200,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter()

			handler := MetricListHandler{
				Service: test.args,
			}
			r.Get("/", handler.ListMetrics)

			ts := httptest.NewServer(r)
			defer ts.Close()

			resp := testRequestMetricListHandler(t, ts, test.method, test.url)
			defer resp.Body.Close()

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}

func testRequestMetricListHandler(t *testing.T, ts *httptest.Server, method, path string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	return resp
}