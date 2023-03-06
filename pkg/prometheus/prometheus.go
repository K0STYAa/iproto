package prometheus

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const (
	timeoutSec = 3
	latencyStart = 0.001
	latencyWidth = 0.001
	latencyCount = 10
)

var (
	ConnectionsCount = prometheus.NewGauge(prometheus.GaugeOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "active_connections",
		Help: "Total number of connections",
	})

	StorageReads = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_reads_total",
		Help: "Total number of storage reads",
	})

	StorageWrites = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_writes_total",
		Help: "Total number of storage writes",
	})

	SuccessfulStorageReads = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_successful_reads_total",
		Help: "Total number of successful storage reads",
	})

	SuccessfulStorageWrites = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_successful_writes_total",
		Help: "Total number of successful storage writes",
	})

	ReadWiriteChangeState = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "change_storage_state_on_read_write_total",
		Help: "Total number of ReadWirite changing storage state",
	})

	ReadOnlyChangeState = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "change_storage_state_on_read_only_total",
		Help: "Total number of ReadOnly changing storage state",
	})

	MaintenanceChangeState = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "change_storage_state_on_maintenance_total",
		Help: "Total number of Maintenance changing storage state",
	})

	ErrorStorageReads = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_error_reads_total",
		Help: "Total number of errors at storage reads",
	})

	ErrorStorageWrites = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: lll,exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_error_writes_total",
		Help: "Total number of errors at storage writes",
	})

	Latency = prometheus.NewHistogram(prometheus.HistogramOpts{ //nolint: exhaustivestruct,exhaustruct,gochecknoglobals
		Name:    "RPC_request_duration_seconds",
		Help:    "Histogram of RPC request latencies",
		Buckets: prometheus.LinearBuckets(latencyStart, latencyWidth, latencyCount),
	})
)

func InitPrometheus() {
	prometheus.MustRegister(StorageReads)
	prometheus.MustRegister(StorageWrites)
	prometheus.MustRegister(SuccessfulStorageReads)
	prometheus.MustRegister(SuccessfulStorageWrites)
	prometheus.MustRegister(ErrorStorageReads)
	prometheus.MustRegister(ErrorStorageWrites)
	prometheus.MustRegister(ReadWiriteChangeState)
	prometheus.MustRegister(ReadOnlyChangeState)
	prometheus.MustRegister(MaintenanceChangeState)
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