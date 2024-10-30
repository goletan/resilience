// /resilience/retry/metrics.go
package retry

import (
	"github.com/goletan/observability/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type RetryMetrics struct{}

// Retry Metrics: Track retry attempts.
var (
	RetryAttempts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "goletan",
			Subsystem: "resilience",
			Name:      "retry_attempts_total",
			Help:      "Counts the number of retry attempts for operations.",
		},
		[]string{"operation", "status"},
	)
)

func InitMetrics() {
	metrics.NewManager().Register(&RetryMetrics{})
}

func (rm *RetryMetrics) Register() error {
	if err := prometheus.Register(RetryAttempts); err != nil {
		return err
	}

	return nil
}

// CountRetryAttempt logs retry attempts for operations.
func CountRetryAttempt(operation, status string) {
	RetryAttempts.WithLabelValues(operation, status).Inc()
}
