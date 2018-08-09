package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"os"
	"time"
)

func init() {
	// Metrics
	metricsRegistry := metrics.NewRegistry()
	metrics.RegisterDebugGCStats(metricsRegistry)
	metrics.RegisterRuntimeMemStats(metricsRegistry)

	go metrics.CaptureDebugGCStats(metricsRegistry, time.Second*5)
	go metrics.CaptureRuntimeMemStats(metricsRegistry, time.Second*5)
	go influxdb.InfluxDB(
		metricsRegistry,
		time.Second*5,
		"http://localhost:8086",
		"metrics",
		"",
		"",
	)

	// Clog
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100,
	}); err != nil {
		fmt.Printf("[INFO] Init console log failed. error %+v.", err)
		os.Exit(1)
	}

	if err := clog.New(clog.FILE, clog.FileConfig{
		Level:      clog.INFO,
		BufferSize: 100,
		Filename:   "logs/clog.log",
		FileRotationConfig: clog.FileRotationConfig{
			Rotate: true,
			Daily:  true,
		},
	}); err != nil {
		fmt.Printf("[INFO] Init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

func main() {
	defer clog.Shutdown()

	node, err := snowflake.NewNode(1)
	exceptions.CheckError(err)

	requestID := node.Generate()
	clog.Info("Request_id: %s", requestID.Base58())

	cfg, err := config.Init()
	exceptions.CheckError(err)

	http := admin.HTTP{Config: &cfg}
	http.Init()
}
