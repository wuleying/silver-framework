package metrics

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
	"time"
)

type Metric struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (m *Metric) Init() {
	metricsRegistry := metrics.NewRegistry()
	metrics.RegisterDebugGCStats(metricsRegistry)
	metrics.RegisterRuntimeMemStats(metricsRegistry)

	go metrics.CaptureDebugGCStats(metricsRegistry, time.Second*5)
	go metrics.CaptureRuntimeMemStats(metricsRegistry, time.Second*5)

	go influxdb.InfluxDB(
		metricsRegistry,
		time.Second*5,
		fmt.Sprintf("http://%s:%s", m.Host, m.Port),
		m.Database,
		m.Username,
		m.Password,
	)
}
