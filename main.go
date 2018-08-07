package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"os"
)

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
	defer clog.Shutdown()

	node, err := snowflake.NewNode(1)
	exceptions.CheckError(err)

	requestId := node.Generate()
	clog.Info("Time: %d", requestId.Time())
	clog.Info("Node: %d", requestId.Node())
	clog.Info("Step: %d", requestId.Step())
	clog.Info("Request_id: %s", requestId.Base58())

	cfg, err := config.Init()
	exceptions.CheckError(err)

	http := admin.HTTP{Config: &cfg}
	http.Init()
}
