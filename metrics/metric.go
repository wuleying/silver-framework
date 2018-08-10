package metrics

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
	"github.com/wuleying/silver-framework/config"
	"time"
)

type Metric struct {
	Config *config.Config
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
		fmt.Sprintf("http://%s:%s", m.Config.Metrics["host"], m.Config.Metrics["port"]),
		m.Config.Metrics["database"],
		m.Config.Metrics["username"],
		m.Config.Metrics["password"],
	)
}
