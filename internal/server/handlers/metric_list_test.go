package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sur1k1/go-metrics/internal/server/storage"
)

func TestMetricListHandler(t *testing.T) {
	tests := []struct {
		name       string
		args       AllMetricsGetter
		method		string
		url				string
		wantStatus int
		wantBody   string
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
			wantBody: "Metrics List:\nalloc: 123.000\npollcount: 123\n",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/", MetricListHandler(test.args))

			ts := httptest.NewServer(r)
			defer ts.Close()

			resp := testRequestMetricListHandler(t, ts, test.method, test.url)
			defer resp.Body.Close()

			respBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, test.wantStatus, resp.StatusCode)
			assert.Equal(t, test.wantBody, string(respBody))
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