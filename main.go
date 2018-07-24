package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"os"
	"time"
	"math/rand"

	appm "github.com/wgliang/goappmonitor"
)

// Base or system performance data,such as memeory,gc,network and so on.
func baseOrsystem() {
	for _ = range time.Tick(time.Second*time.Duration(10)) {
		// (commonly used) Meter, used to sum and calculate the rate of change. Use scenarios
		// such as the number of home visits statistics, CG etc..
		pv := int64(rand.Int31n(100))
		appm.Meter("appm.meter", pv)
		appm.Meter("appm.meter.2", pv-50)

		// (commonly used) Gauge, used to preserve the value of the instantaneous value of the
		// type of record. Use scenarios such as statistical queue length, statistics CPU usage,
		// and so on.
		queueSize := int64(rand.Int31n(100) - 50)
		appm.Gauge("appm.gauge", queueSize)

		cpuUtil := float64(rand.Int31n(10000)) / float64(100)
		appm.GaugeFloat64("appm.gauge.float64", cpuUtil)
	}
}

// Custom or business performance data,such as qps,num of function be called, task queue and so on.
func customOrbusiness() {
	for _ = range time.Tick(time.Second) {
		// Histogram, using the exponential decay sampling method, the probability distribution of
		// the statistical object is calculated. Using scenarios such as the probability distribution
		// of the statistics home page to access the delay
		delay := int64(rand.Int31n(100))
		appm.Histogram("appm.histogram", delay)
	}
}

func init() {
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
	var ch chan int
	go baseOrsystem()
	go customOrbusiness()
	<-ch

	defer clog.Shutdown()

	cfg, err := config.Init()
	exceptions.CheckError(err)

	h := admin.HTTP{Config: &cfg}
	h.Init()
}
