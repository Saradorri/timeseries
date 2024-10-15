package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
)

var (
	// QueriesTotal Total number of queries
	QueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "queries_total",
			Help: "Total number of queries made to Prometheus",
		},
		[]string{"status"},
	)

	// QueryDuration measure query response time
	QueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "query_duration_seconds",
			Help:    "Duration of queries to Prometheus in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	// TimeSeriesValue get recorded values at a specific timestamp
	TimeSeriesValue = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "time_series_value",
		Help: "Recorded values at a specific timestamp.",
	},
		[]string{"timestamp"},
	)
)

func init() {
	prometheus.MustRegister(QueriesTotal)
	prometheus.MustRegister(QueryDuration)
	prometheus.MustRegister(TimeSeriesValue)
}

func StartServer(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(strconv.Itoa(port), nil); err != nil {
			log.Fatalf("Error starting metrics server: %v", err)
		}
		log.Printf("Metrics server started on port %d", port)
	}()
}

func SetValue(timestamp string, value float64) {
	TimeSeriesValue.With(prometheus.Labels{"timestamp": timestamp}).Set(value)
}
