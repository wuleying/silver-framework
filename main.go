package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/metrics"
	"os"
)

// initClog
func initClog() {
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

// initRequestID

func main() {
	// Clog
	initClog()
	defer clog.Shutdown()

	cfg, err := config.Init()
	exceptions.CheckError(err)

	metric := metrics.Metric{
		Host:     cfg["metric"]["host"],
		Port:     cfg["metric"]["port"],
		Database: cfg["metric"]["database"],
		Username: cfg["metric"]["username"],
		Password: cfg["metric"]["password"],
	}
	metric.Init()

	http := admin.HTTP{Config: cfg}
	http.Init()
}
