package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddGauge(t *testing.T) {
	type metric struct {
		name  string
		value string
	}

	tests := []struct {
		name    string
		storage *MemStorage
		metric  metric
		wantErr bool
	}{
		{
			name: "simple test #1",
			storage: &MemStorage{
				GaugeMap: map[string]float64{
					"alloc": 209.0,
				},
			},
			metric: metric{
				name:  "alloc",
				value: "203.0",
			},
			wantErr: false,
		},
		{
			name:    "big float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "922337203685477580000007.0",
			},
			wantErr: false,
		},
		{
			name:    "big int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "9223372036854775800000070",
			},
			wantErr: false,
		},
		{
			name:    "negative big float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "-922337203685477580000007.0",
			},
			wantErr: false,
		},
		{
			name:    "negative big int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "-9223372036854775800000070",
			},
			wantErr: false,
		},
		{
			name:    "zero int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "0",
			},
			wantErr: false,
		},
		{
			name:    "zero float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "0.0",
			},
			wantErr: false,
		},
		{
			name:    "negative zero int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "-0",
			},
			wantErr: false,
		},
		{
			name:    "negative zero float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "-0.0",
			},
			wantErr: false,
		},
		{
			name:    "string value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name:  "alloc",
				value: "hello world!",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.storage.AddGauge(test.metric.name, test.metric.value)
			if !test.wantErr {
				require.NoError(t, err)

				assert.Contains(t, test.storage.GaugeMap, test.metric.name)
				return
			}
			assert.Error(t, err)
		})
	}
}

func TestAddCounter(t *testing.T) {
	type metric struct {
		name  string
		value string
	}

	tests := []struct {
		name      string
		storage   *MemStorage
		metric    metric
		wantErr   bool
		wantValue int64
	}{
		{
			name: "simple test #1",
			storage: &MemStorage{
				CounterMap: map[string]int64{
					"pollcount": 203,
				},
			},
			metric: metric{
				name:  "pollcount",
				value: "203",
			},
			wantErr:   false,
			wantValue: 406,
		},
		{
			name: "big float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 92233720368547758,
			}},
			metric: metric{
				name:  "pollcount",
				value: "92233720368547758.6",
			},
			wantErr:   true,
			wantValue: int64(2 * 92233720368547758),
		},
		{
			name: "big int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 922337203685477580,
			}},
			metric: metric{
				name:  "pollcount",
				value: "922337203685477580",
			},
			wantErr:   false,
			wantValue: int64(2 * 922337203685477580),
		},
		{
			name: "negative big float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 9223372036854775807,
			}},
			metric: metric{
				name:  "pollcount",
				value: "-9223372036854775807.9",
			},
			wantErr:   true,
			wantValue: 0,
		},
		{
			name: "negative big int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 9223372036854775807,
			}},
			metric: metric{
				name:  "pollcount",
				value: "-9223372036854775807",
			},
			wantErr:   false,
			wantValue: 0,
		},
		{
			name: "zero int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 5,
			}},
			metric: metric{
				name:  "pollcount",
				value: "0",
			},
			wantErr:   false,
			wantValue: 5,
		},
		{
			name: "zero float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 9223372036854775807,
			}},
			metric: metric{
				name:  "pollcount",
				value: "0.8",
			},
			wantErr:   true,
			wantValue: 9223372036854775807,
		},
		{
			name: "negative zero int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 6,
			}},
			metric: metric{
				name:  "pollcount",
				value: "-0",
			},
			wantErr:   false,
			wantValue: 6,
		},
		{
			name: "negative zero float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"pollcount": 0,
			}},
			metric: metric{
				name:  "pollcount",
				value: "-0.2",
			},
			wantErr:   true,
			wantValue: 0,
		},
		{
			name:    "string value test",
			storage: &MemStorage{CounterMap: map[string]int64{}},
			metric: metric{
				name:  "pollcount",
				value: "hello world!",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.storage.AddCounter(test.metric.name, test.metric.value)
			if !test.wantErr {
				require.NoError(t, err)

				assert.Equal(t, test.wantValue, test.storage.CounterMap[test.metric.name])
				return
			}
			assert.Error(t, err)
		})
	}
}

func TestGetMetric(t *testing.T) {
	type metric struct {
		metricType, metricName string
	}

	tests := []struct {
		name    string
		storage *MemStorage
		metric  metric
		wantErr bool
		want    string
	}{
		{
			name: "get counter metric",
			storage: &MemStorage{
				CounterMap: map[string]int64{
					"pollcount": 123,
				},
			},
			metric: metric{
				metricType: "counter",
				metricName: "pollcount",
			},
			want:    "123",
			wantErr: false,
		},
		{
			name: "get gauge metric",
			storage: &MemStorage{
				GaugeMap: map[string]float64{
					"alloc": 123.0,
				},
			},
			metric: metric{
				metricType: "gauge",
				metricName: "alloc",
			},
			want:    "123.000",
			wantErr: false,
		},
		{
			name: "get metric with invalid type",
			storage: &MemStorage{
				GaugeMap: map[string]float64{},
			},
			metric: metric{
				metricType: "gofer go",
				metricName: "alloc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "get metric with invalid name",
			storage: &MemStorage{
				GaugeMap: map[string]float64{},
			},
			metric: metric{
				metricType: "",
				metricName: "A)_S))S)D)AS)D)",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "get metric with CAPS name",
			storage: &MemStorage{
				GaugeMap: map[string]float64{
					"alloc": 123.0,
				},
			},
			metric: metric{
				metricType: "gauge",
				metricName: "alloc",
			},
			want:    "123.000",
			wantErr: false,
		},
		// {
		// 	name: "get metric with CAPS type",
		// 	storage: &MemStorage{
		// 		GaugeMap: map[string]float64{
		// 			"alloc": 123.0,
		// 		},
		// 	},
		// 	metric: metric{
		// 		metricType: "GAUGE",
		// 		metricName: "alloc",
		// 	},
		// 	want: "123.000000",
		// 	wantErr: false,
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := test.storage.GetMetric(test.metric.metricType, test.metric.metricName)
			if !test.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, test.want, value)
				return
			}
			assert.Error(t, err)
		})
	}
}

func TestGetAllMetrics(t *testing.T) {
	tests := []struct {
		name string
		s    *MemStorage
		want map[string]string
	}{
		{
			name: "positive test #1",
			s: &MemStorage{
				GaugeMap: map[string]float64{
					"alloc": 123.0,
				},
				CounterMap: map[string]int64{
					"pollcount": 123,
				},
			},
			want: map[string]string{
				"alloc": "123.000",
				"pollcount": "123",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			metrics := test.s.GetAllMetrics()
			assert.Equal(t, test.want, metrics)
		})
	}
}
