package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddGauge(t *testing.T) {
	type metric struct{
		name string
		value string
	}

	tests := []struct{
		name string
		storage *MemStorage
		metric metric
		wantErr bool
	}{
		{
			name: "simple test #1",
			storage: &MemStorage{
				GaugeMap: map[string]float64{
					"Alloc":209.0,
				},
			},
			metric: metric{
				name: "Alloc",
				value: "203.0",
			},
			wantErr: false,
		},
		{
			name: "big float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "922337203685477580000007.0",
			},
			wantErr: false,
		},
		{
			name: "big int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "9223372036854775800000070",
			},
			wantErr: false,
		},
		{
			name: "negative big float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "-922337203685477580000007.0",
			},
			wantErr: false,
		},
		{
			name: "negative big int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "-9223372036854775800000070",
			},
			wantErr: false,
		},
		{
			name: "zero int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "0",
			},
			wantErr: false,
		},
		{
			name: "zero float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "0.0",
			},
			wantErr: false,
		},
		{
			name: "negative zero int value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "-0",
			},
			wantErr: false,
		},
		{
			name: "negative zero float value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "-0.0",
			},
			wantErr: false,
		},
		{
			name: "string value test",
			storage: &MemStorage{GaugeMap: map[string]float64{}},
			metric: metric{
				name: "Alloc",
				value: "hello world!",
			},
			wantErr: true,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			err := test.storage.AddGauge(test.metric.name, test.metric.value)
			if !test.wantErr{
				require.NoError(t, err)

				assert.Contains(t, test.storage.GaugeMap, test.metric.name)
				return
			}
			assert.Error(t, err)
		})
	}
}

func TestAddCounter(t *testing.T) {
	type metric struct{
		name string
		value string
	}

	tests := []struct{
		name string
		storage *MemStorage
		metric metric
		wantErr bool
		wantValue int64
	}{
		{
			name: "simple test #1",
			storage: &MemStorage{
				CounterMap: map[string]int64{
					"PollCount": 203,
				},
			},
			metric: metric{
				name: "PollCount",
				value: "203",
			},
			wantErr: false,
			wantValue: 406,
		},
		{
			name: "big float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 92233720368547758,
			}},
			metric: metric{
				name: "PollCount",
				value: "92233720368547758.6",
			},
			wantErr: true,
			wantValue: int64(2*92233720368547758),
		},
		{
			name: "big int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 922337203685477580,
			}},
			metric: metric{
				name: "PollCount",
				value: "922337203685477580",
			},
			wantErr: false,
			wantValue: int64(2*922337203685477580),
		},
		{
			name: "negative big float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 9223372036854775807,
			}},
			metric: metric{
				name: "PollCount",
				value: "-9223372036854775807.9",
			},
			wantErr: true,
			wantValue: 0,
		},
		{
			name: "negative big int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 9223372036854775807,
			}},
			metric: metric{
				name: "PollCount",
				value: "-9223372036854775807",
			},
			wantErr: false,
			wantValue: 0,
		},
		{
			name: "zero int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 5,
			}},
			metric: metric{
				name: "PollCount",
				value: "0",
			},
			wantErr: false,
			wantValue: 5,
		},
		{
			name: "zero float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 9223372036854775807,
			}},
			metric: metric{
				name: "PollCount",
				value: "0.8",
			},
			wantErr: true,
			wantValue: 9223372036854775807,
		},
		{
			name: "negative zero int value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 6,
			}},
			metric: metric{
				name: "PollCount",
				value: "-0",
			},
			wantErr: false,
			wantValue: 6,
		},
		{
			name: "negative zero float value test",
			storage: &MemStorage{CounterMap: map[string]int64{
				"PollCount": 0,
			}},
			metric: metric{
				name: "PollCount",
				value: "-0.2",
			},
			wantErr: true,
			wantValue: 0,
		},
		{
			name: "string value test",
			storage: &MemStorage{CounterMap: map[string]int64{}},
			metric: metric{
				name: "PollCount",
				value: "hello world!",
			},
			wantErr: true,
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			err := test.storage.AddCounter(test.metric.name, test.metric.value)
			if !test.wantErr{
				require.NoError(t, err)

				assert.Equal(t, test.wantValue, test.storage.CounterMap[test.metric.name])
				return
			}
			assert.Error(t, err)
		})
	}
}