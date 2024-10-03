package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sur1k1/go-metrics/internal/server/storage"
)

func TestUpdateHandler(t *testing.T) {
	type args struct {
		s Storage
	}

	type header struct {
		key, value string
	}
	
	tests := []struct {
		name 				string
		args 				args
		url 				string
		header 			header
		method			string
		wantStatus 	int
	}{
		{
			name: "simple gauge test #1",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/gauge/alloc/123.0",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 200,
		},
		{
			name: "simple counter test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/counter/pollCount/123",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 200,
		},
		{
			name: "invalid content type test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/counter/pollCount/123",
			header: header{
				key: "Content-Type",
				value: "text/html",
			},
			method: "POST",
			wantStatus: 415,
		},
		{
			name: "invalid requst method (GET) test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/counter/pollCount/123",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "GET",
			wantStatus: 405,
		},
		{
			name: "invalid metric name test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/counter//123",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 404,
		},
		{
			name: "invalid value test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/counter/pollCount/hello_world",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 400,
		},
		{
			name: "invalid type metric name test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/aasdkok/alloc/123",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 400,
		},
		{
			name: "invalid url test",
			args: args{
				s: &storage.MemStorage{
					GaugeMap: map[string]float64{},
					CounterMap: map[string]int64{},
				},
			},
			url: "/update/hi",
			header: header{
				key: "Content-Type",
				value: "text/plain",
			},
			method: "POST",
			wantStatus: 404,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/update/{type}/{metric}/{value}", UpdateHandler(test.args.s))

			ts := httptest.NewServer(r)
			defer ts.Close()
			
			resp := testRequest_UpdateHandler(t, ts, test.method, test.url, test.header.key, test.header.value)
			defer resp.Body.Close()

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}

func testRequest_UpdateHandler(t *testing.T, ts *httptest.Server, method, path, headerKey, headerValue string) *http.Response {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	req.Header.Set(headerKey, headerValue)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	return resp
}