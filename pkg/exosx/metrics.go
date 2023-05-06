package exosx

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	APICallMetric = "exosx_api_appliance_api_call"
	APICallHelp   = "How many API calls have been executed"

	APICallDurationMetric = "exosx_api_appliance_api_call_duration"
	APICallDurationHelp   = "The total duration of API calls"
)

type Collector struct {
	apiCall         *prometheus.CounterVec
	apiCallDuration *prometheus.CounterVec
}

func newCollector() *Collector {
	return &Collector{
		apiCall: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: APICallMetric,
				Help: APICallHelp,
			},
			[]string{"endpoint", "success"},
		),
		apiCallDuration: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: APICallDurationMetric,
				Help: APICallDurationHelp,
			},
			[]string{"endpoint"},
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	collector.apiCall.Describe(ch)
	collector.apiCallDuration.Describe(ch)
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	collector.apiCall.Collect(ch)
	collector.apiCallDuration.Collect(ch)
}

func (collector *Collector) trackAPICall(endpoint string) func(bool) {
	start := time.Now()

	return func(success bool) {
		duration := float64(time.Since(start).Nanoseconds()) / 1000 / 1000 / 1000
		collector.apiCallDuration.WithLabelValues(endpoint).Add(duration)
		collector.apiCall.WithLabelValues(endpoint, fmt.Sprintf("%t", success)).Inc()
	}
}
