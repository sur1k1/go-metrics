package metric

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"runtime"
	"time"
)

const (
	serverAddress		string = "localhost:8080"
	reportForm			string = "http://%s/update/%s/%s/%v"
	typeGauge				string = "gauge"
	typeCounter			string = "counter"
	pollInterval   	time.Duration = 2
	reportInterval 	time.Duration = 10
)

// Временное хранилище метрик
type MetricStorage struct {
	GaugeMap   map[string]float64
	CounterMap map[string]int64
}

// Получение MetricStorage клиента
func NewMetricStorage() *MetricStorage {
	return &MetricStorage{
		GaugeMap: map[string]float64{},
		CounterMap: map[string]int64{},
	}
}

// Описание интерфейса метрик
type Metric interface {
	MetricUpdaterIntf
	MetricSenderIntf
}

// Интерфейс обновления метрик
type MetricUpdaterIntf interface {
	UpdateCounter()
	UpdateGauge()
}

// Получение обновление по pollInterval
func MetricUpdater(mu MetricUpdaterIntf) {
	for {
		mu.UpdateCounter()
		mu.UpdateGauge()
		time.Sleep(time.Second*pollInterval)
	}
}

// Обновление gauge метрик
func (m *MetricStorage) UpdateGauge() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	m.GaugeMap["Alloc"] = float64(rtm.Alloc)
	m.GaugeMap["BuckHashSys"] = float64(rtm.BuckHashSys)
	m.GaugeMap["Frees"] = float64(rtm.Frees)
	m.GaugeMap["GCCPUFraction"] = float64(rtm.GCCPUFraction)
	m.GaugeMap["GCSys"] = float64(rtm.GCSys)
	m.GaugeMap["HeapAlloc"] = float64(rtm.HeapAlloc)
	m.GaugeMap["HeapIdle"] = float64(rtm.HeapIdle)
	m.GaugeMap["HeapInuse"] = float64(rtm.HeapInuse)
	m.GaugeMap["HeapObjects"] = float64(rtm.HeapObjects)
	m.GaugeMap["HeapReleased"] = float64(rtm.HeapReleased)
	m.GaugeMap["HeapSys"] = float64(rtm.HeapSys)
	m.GaugeMap["LastGC"] = float64(rtm.LastGC)
	m.GaugeMap["Lookups"] = float64(rtm.Lookups)
	m.GaugeMap["MCacheInuse"] = float64(rtm.MCacheInuse)
	m.GaugeMap["MCacheSys"] = float64(rtm.MCacheSys)
	m.GaugeMap["MSpanInuse"] = float64(rtm.MSpanInuse)
	m.GaugeMap["MSpanSys"] = float64(rtm.MSpanSys)
	m.GaugeMap["Mallocs"] = float64(rtm.Mallocs)
	m.GaugeMap["NextGC"] = float64(rtm.NextGC)
	m.GaugeMap["NumForcedGC"] = float64(rtm.NumForcedGC)
	m.GaugeMap["NumGC"] = float64(rtm.NumGC)
	m.GaugeMap["OtherSys"] = float64(rtm.OtherSys)
	m.GaugeMap["PauseTotalNs"] = float64(rtm.PauseTotalNs)
	m.GaugeMap["StackInuse"] = float64(rtm.StackInuse)
	m.GaugeMap["StackSys"] = float64(rtm.StackSys)
	m.GaugeMap["Sys"] = float64(rtm.Sys)
	m.GaugeMap["TotalAlloc"] = float64(rtm.TotalAlloc)
	m.GaugeMap["RandomValue"] = rand.Float64()
}

// Обновление counter метрик
func (m *MetricStorage) UpdateCounter() {
	m.CounterMap["PollCount"] += 1
}

// Интерфейс отправки метрик
type MetricSenderIntf interface {
	Send(client http.Client) error
}

// Отправка метрик на сервер по reportInterval
func MetricSender(ms MetricSenderIntf) {
	var client http.Client
	for {
		err := ms.Send(client)
		if err != nil{
			log.Println(err)
			continue
		}

		time.Sleep(time.Second*reportInterval)
	}
}

// Фнкция отправки метрик на сервер
func (m *MetricStorage) Send(client http.Client) error {
	for name, value := range m.GaugeMap{
		url := fmt.Sprintf(reportForm, serverAddress, typeGauge, name, value)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil{
			return err
		}
		req.Header.Set("Content-Type", "text/plain")

		resp, err := client.Do(req)
		if err != nil{
			return err
		}
		defer resp.Body.Close()
		
		_, err = io.ReadAll(resp.Body)
		if err != nil{
			return err
		}
	}

	for name, value := range m.CounterMap{
		url := fmt.Sprintf(reportForm, serverAddress, typeCounter, name, value)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil{
			return err
		}
		req.Header.Set("Content-Type", "text/plain")

		resp, err := client.Do(req)
		if err != nil{
			return err
		}
		defer resp.Body.Close()
		
		_, err = io.ReadAll(resp.Body)
		if err != nil{
			return err
		}
		req.Header.Set("Content-Type", "text/plain")
	}

	return nil
}