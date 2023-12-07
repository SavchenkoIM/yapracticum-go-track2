package metricsPoll

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
)

type metricsData struct {
	name  string
	typ   string
	value string
}

type MetricsHandler struct {
	metricsSlice []metricsData
	counter      int64
	client       http.Client
}

func NewMetricsHandler() MetricsHandler {
	return MetricsHandler{
		metricsSlice: []metricsData{
			{name: "Alloc", typ: "gauge", value: ""},
			{name: "BuckHashSys", typ: "gauge", value: ""},
			{name: "Frees", typ: "gauge", value: ""},
			{name: "GCCPUFraction", typ: "gauge", value: ""},
			{name: "GCSys", typ: "gauge", value: ""},
			{name: "HeapAlloc", typ: "gauge", value: ""},
			{name: "HeapIdle", typ: "gauge", value: ""},
			{name: "HeapInuse", typ: "gauge", value: ""},
			{name: "HeapObjects", typ: "gauge", value: ""},
			{name: "HeapReleased", typ: "gauge", value: ""},
			{name: "HeapSys", typ: "gauge", value: ""},
			{name: "LastGC", typ: "gauge", value: ""},
			{name: "Lookups", typ: "gauge", value: ""},
			{name: "MCacheInuse", typ: "gauge", value: ""},
			{name: "MCacheSys", typ: "gauge", value: ""},
			{name: "MSpanInuse", typ: "gauge", value: ""},
			{name: "MSpanSys", typ: "gauge", value: ""},
			{name: "Mallocs", typ: "gauge", value: ""},
			{name: "NextGC", typ: "gauge", value: ""},
			{name: "NumForcedGC", typ: "gauge", value: ""},
			{name: "NumGC", typ: "gauge", value: ""},
			{name: "OtherSys", typ: "gauge", value: ""},
			{name: "PauseTotalNs", typ: "gauge", value: ""},
			{name: "StackInuse", typ: "gauge", value: ""},
			{name: "StackSys", typ: "gauge", value: ""},
			{name: "Sys", typ: "gauge", value: ""},
			{name: "TotalAlloc", typ: "gauge", value: ""},
			{name: "RandomValue", typ: "gauge", value: ""},
			//{name: "PollCount", typ: "counter", value: ""},
		},
	}
}

func (this MetricsHandler) SendData() {
	for _, v := range this.metricsSlice {
		_, err := this.client.Post("http://localhost:8080/update/"+v.typ+"/"+v.name+"/"+v.value,
			"text/plain",
			nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func (this *MetricsHandler) RefreshData() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	this.metricsSlice[0].value = fmt.Sprintf("%v", ms.Alloc)
	this.metricsSlice[1].value = fmt.Sprintf("%v", ms.BuckHashSys)
	this.metricsSlice[2].value = fmt.Sprintf("%v", ms.Frees)
	this.metricsSlice[3].value = fmt.Sprintf("%v", ms.GCCPUFraction)
	this.metricsSlice[4].value = fmt.Sprintf("%v", ms.GCSys)
	this.metricsSlice[5].value = fmt.Sprintf("%v", ms.HeapAlloc)
	this.metricsSlice[6].value = fmt.Sprintf("%v", ms.HeapIdle)
	this.metricsSlice[7].value = fmt.Sprintf("%v", ms.HeapInuse)
	this.metricsSlice[8].value = fmt.Sprintf("%v", ms.HeapObjects)
	this.metricsSlice[9].value = fmt.Sprintf("%v", ms.HeapReleased)
	this.metricsSlice[10].value = fmt.Sprintf("%v", ms.HeapSys)
	this.metricsSlice[11].value = fmt.Sprintf("%v", ms.LastGC)
	this.metricsSlice[12].value = fmt.Sprintf("%v", ms.Lookups)
	this.metricsSlice[13].value = fmt.Sprintf("%v", ms.MCacheInuse)
	this.metricsSlice[14].value = fmt.Sprintf("%v", ms.MCacheSys)
	this.metricsSlice[15].value = fmt.Sprintf("%v", ms.MSpanInuse)
	this.metricsSlice[16].value = fmt.Sprintf("%v", ms.MSpanSys)
	this.metricsSlice[17].value = fmt.Sprintf("%v", ms.Mallocs)
	this.metricsSlice[18].value = fmt.Sprintf("%v", ms.NextGC)
	this.metricsSlice[19].value = fmt.Sprintf("%v", ms.NumForcedGC)
	this.metricsSlice[20].value = fmt.Sprintf("%v", ms.NumGC)
	this.metricsSlice[21].value = fmt.Sprintf("%v", ms.OtherSys)
	this.metricsSlice[22].value = fmt.Sprintf("%v", ms.PauseTotalNs)
	this.metricsSlice[23].value = fmt.Sprintf("%v", ms.StackInuse)
	this.metricsSlice[24].value = fmt.Sprintf("%v", ms.StackSys)
	this.metricsSlice[25].value = fmt.Sprintf("%v", ms.Sys)
	this.metricsSlice[26].value = fmt.Sprintf("%v", ms.TotalAlloc)
	this.metricsSlice[27].value = fmt.Sprintf("%v", rand.Float64())
	//this.metricsSlice[28].value = fmt.Sprintf("%d", this.counter)

	_, err := this.client.Post("http://localhost:8080/update/counter/PollCount/1",
		"text/plain",
		nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}