package metrics

import "github.com/prometheus/client_golang/prometheus"

var methodsCalls = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "methods_calls",
		Help: "Number of service methods calls",
	},
	[]string{"handler"},
)

var consumedMessages = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "consumed_messages",
		Help: "Number of consumed messages",
	},
)

// RegisterMetrics creates metrics handlers
func RegisterMetrics() {
	prometheus.MustRegister(methodsCalls)
	prometheus.MustRegister(consumedMessages)
}

// IncMethodsCalls increments number of service method calls
func IncMethodsCalls(handler string) {
	methodsCalls.With(prometheus.Labels{"handler": handler}).Inc()
}

// IncConsumedMessages increments number of consumed messages
func IncConsumedMessages() {
	consumedMessages.Inc()
}
