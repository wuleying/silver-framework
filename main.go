package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/metrics"
	"os"
)

var requestID snowflake.ID

func init() {
	// Clog
	initClog()

	// Request id
	initRequestID()
}

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
func initRequestID() {
	node, err := snowflake.NewNode(1)
	exceptions.CheckError(err)

	requestID = node.Generate()
	clog.Info("Request_id: %s", requestID.Base58())
}

func main() {
	defer clog.Shutdown()

	cfg, err := config.Init()
	exceptions.CheckError(err)

	metric := metrics.Metric{Config: &cfg}
	metric.Init()

	http := admin.HTTP{Config: &cfg}
	http.Init()
}
