package metrics

import (
	echoprom "github.com/labstack/echo-contrib/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	PipelinesStartedCount *echoprom.Metric
	PipelinesDuration     *echoprom.Metric
}

func (m *Metrics) MetricList() []*echoprom.Metric {
	return []*echoprom.Metric{
		m.PipelinesStartedCount,
		m.PipelinesDuration,
	}
}

// counter, counter_vec, gauge, gauge_vec,
// histogram, histogram_vec, summary, summary_vec
func NewMetrics() *Metrics {
	return &Metrics{
		PipelinesStartedCount: &echoprom.Metric{
			Name:        "pipelines_started_count",
			Description: "Count of pipelines started",
			Type:        "counter_vec",
			Args:        []string{"project", "source", "ref"},
		},
		PipelinesDuration: &echoprom.Metric{
			Name:        "pipelines_duration_seconds",
			Description: "Pipelines duration in seconds",
			Type:        "histogram_vec",
			Args:        []string{"project", "source", "ref", "status"},
			Buckets:     prometheus.LinearBuckets(30, 30, 40),
		},
	}
}

func CounterIncrease(metric *echoprom.Metric, labels map[string]string) {
	metric.MetricCollector.(*prometheus.CounterVec).With(labels).Inc()
}

func HistogramObserve(metric *echoprom.Metric, observedValue float64, labels map[string]string) {
	metric.MetricCollector.(*prometheus.HistogramVec).With(labels).Observe(observedValue)
}
