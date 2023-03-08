package prometheus

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const (
	timeoutSec   = 3
	latencyStart = 0.001
	latencyWidth = 0.001
	latencyCount = 10
)

var (
	ConnectionsCount = prometheus.NewGauge(//nolint: gochecknoglobals
		prometheus.GaugeOpts{ //nolint: exhaustivestruct,exhaustruct
			Name: "active_connections",
			Help: "Total number of connections",
		},
	)

	ApiCall = prometheus.NewCounterVec( //nolint: exhaustivestruct,exhaustruct,gochecknoglobals
		prometheus.CounterOpts{
			Name: "api_call_total",
			Help: "Total number of api calls by methods",
		},
        []string{"method"},
	)

	SuccessfulStorageReadsWrites = prometheus.NewCounterVec( //nolint: gochecknoglobals
		prometheus.CounterOpts{ //nolint: exhaustivestruct,exhaustruct
			Name: "storage_successful_reads_writes_total",
			Help: "Total number of successful storage reads and writes",
		},
        []string{"action"},
	)

	ErrorStorageReadsWrites = prometheus.NewCounterVec( //nolint: gochecknoglobals
		prometheus.CounterOpts{ //nolint: exhaustivestruct,exhaustruct
			Name: "storage_error_reads_writes_total",
			Help: "Total number of errors at storage reads and writes",
		},
        []string{"action"},
	)

	Latency = prometheus.NewHistogram( //nolint: gochecknoglobals
		prometheus.HistogramOpts{ //nolint: exhaustivestruct,exhaustruct
			Name:    "RPC_request_duration_seconds",
			Help:    "Histogram of RPC request latencies",
			Buckets: prometheus.LinearBuckets(latencyStart, latencyWidth, latencyCount),
		},
	)
)

func InitPrometheus() {
	prometheus.MustRegister(ApiCall)
	prometheus.MustRegister(SuccessfulStorageReadsWrites)
	prometheus.MustRegister(ErrorStorageReadsWrites)
	prometheus.MustRegister(ConnectionsCount)
	prometheus.MustRegister(Latency)

	http.Handle("/metrics", promhttp.Handler())

	server := &http.Server{ //nolint: exhaustivestruct,exhaustruct
		Addr:              ":8088",
		ReadHeaderTimeout: timeoutSec * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal("ListenHTTP error: ", err)
	}
}
