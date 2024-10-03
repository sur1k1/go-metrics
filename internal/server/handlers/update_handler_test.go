package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
			ts := httptest.NewServer(UpdateHandler(test.args.s))
			defer ts.Close()
			
			request := httptest.NewRequest(test.method, test.url, nil)
			request.Header.Set(test.header.key, test.header.value)
			w := httptest.NewRecorder()

			handler := UpdateHandler(test.args.s)
			handler(w, request)

			res := w.Result()
			defer res.Body.Close()
			_, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			assert.Equal(t, test.wantStatus, res.StatusCode)
		})
	}
}