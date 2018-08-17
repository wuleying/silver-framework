package main

import (
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/llog"
	"github.com/wuleying/silver-framework/metrics"
)

func main() {
	// Log init
	llog.Init()
	defer clog.Shutdown()

	cfg, err := config.Init()
	exceptions.CheckError(err)

	metric := metrics.Metric{
		Host:     cfg["metrics"]["host"],
		Port:     cfg["metrics"]["port"],
		Database: cfg["metrics"]["database"],
		Username: cfg["metrics"]["username"],
		Password: cfg["metrics"]["password"],
	}
	metric.Init()

	http := admin.HTTP{
		Host: cfg["setting"]["host"],
		Port: cfg["setting"]["port"],
	}
	http.Init()
}
