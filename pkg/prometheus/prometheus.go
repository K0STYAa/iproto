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
)

var (
	StorageReads = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_reads_total",
		Help: "Total number of storage reads",
	})

	StorageWrites = prometheus.NewCounter(prometheus.CounterOpts{ //nolint: exhaustivestruct,exhaustruct,gochecknoglobals
		Name: "storage_writes_total",
		Help: "Total number of storage writes",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(StorageReads)
	prometheus.MustRegister(StorageWrites)

	http.Handle("/metrics", promhttp.Handler())

	server := &http.Server{ //nolint: exhaustivestruct,exhaustruct
		Addr:              ":8088",
		ReadHeaderTimeout: timeoutSec * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal("ListenHTTP error: ", err)
	}
}
