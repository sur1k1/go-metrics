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

func TestMetricHandler(t *testing.T) {
	tests := []struct{
		name string
		args MetricService
		url string
		wantStatus int
		method string
	}{
		{
			name: "positive test #1",
			args: &storage.MemStorage{
				CounterMap: map[string]int64{
					"pollcount": 123,
				},
			},
			method: "GET",
			url: "/value/counter/pollcount",
			wantStatus: 200,
		},
		{
			name: "autotest 404",
			args: &storage.MemStorage{
				CounterMap: map[string]int64{
					"pollcount": 123,
				},
			},
			method: "GET",
			url: "/value/gauge/testSetGet203",
			wantStatus: 404,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter()

			handler := MetricHandler{
				Service: test.args,
			}
			r.Get("/value/{type}/{metric}", handler.MetricValue())

			ts := httptest.NewServer(r)
			defer ts.Close()

			resp := testRequestMetricHandler(t, ts, test.method, test.url)
			defer resp.Body.Close()

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}

func testRequestMetricHandler(t *testing.T, ts *httptest.Server, method, path string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)

	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	return resp
}