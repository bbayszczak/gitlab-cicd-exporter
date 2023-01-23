package metrics

import (
	echoprom "github.com/labstack/echo-contrib/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	PipelinesStartedCount *echoprom.Metric
	PipelinesEndedCount   *echoprom.Metric
}

func (m *Metrics) MetricList() []*echoprom.Metric {
	return []*echoprom.Metric{
		m.PipelinesStartedCount,
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
			Args:        []string{"project"},
		},
	}
}

func CounterIncrease(metric *echoprom.Metric, labels map[string]string) {
	metric.MetricCollector.(*prometheus.CounterVec).With(labels).Inc()
}
