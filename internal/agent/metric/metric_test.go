package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricStorage_UpdateGauge(t *testing.T) {
	gaugeMetric := []string{
		"Alloc", "BuckHashSys", "Frees", "BuckHashSys",
		"GCCPUFraction", "GCSys", "HeapAlloc", "HeapIdle",
		"HeapInuse", "HeapObjects", "HeapReleased", "HeapSys",
		"LastGC", "Lookups", "MCacheInuse", "MCacheSys",
		"MSpanInuse", "MSpanSys", "Mallocs", "NextGC",
		"NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs",
		"StackInuse", "StackSys", "Sys", "TotalAlloc",
		"RandomValue",
	}

	tests := []struct {
		name string
		m    *MetricStorage
	}{
		{
			name: "positive test #1",
			m: &MetricStorage{
				GaugeMap: map[string]float64{},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.m.UpdateGauge()

			for _, m := range gaugeMetric{
				assert.Contains(t, test.m.GaugeMap, m) 
			}
		})
	}
}

func TestMetricStorage_UpdateCounter(t *testing.T) {
	tests := []struct {
		name string
		metricName	string
		m    *MetricStorage
		want int64
	}{
		{
			name: "positive test #1",
			metricName: "PollCount",
			m: &MetricStorage{
				CounterMap: map[string]int64{},
			},
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.m.UpdateCounter()

			assert.Equal(t, test.want, test.m.CounterMap[test.metricName])
		})
	}
}

// func TestMetricStorage_Send(t *testing.T) {
// 	type args struct {
// 		client http.Client
// 	}
// 	tests := []struct {
// 		name    string
// 		m       *MetricStorage
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "positive test #1",
// 			m: &MetricStorage{
// 				GaugeMap: map[string]float64{
// 					"Alloc": 123.0,
// 				},
// 				CounterMap: map[string]int64{
// 					"PollCount": 24,
// 				},
// 			},
// 			args: args{client: http.Client{}},
// 			wantErr: false,
// 		},
// 		{
// 			name: "nil test",
// 			m: &MetricStorage{
// 				GaugeMap: map[string]float64{},
// 				CounterMap: map[string]int64{},
// 			},
// 			args: args{client: http.Client{}},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.m.Send(tt.args.client); (err != nil) != tt.wantErr {
// 				t.Errorf("MetricStorage.Send() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
