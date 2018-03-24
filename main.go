package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/globals"
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
}

func main() {
	defer clog.Shutdown()

	config, err := config.Init()
	if err != nil {
		clog.Fatal(globals.ClogSkipDisplayInfo, "Init config failed: %s", err.Error())
	}

	clog.Info("Hello, %s.", config.Setting["project_name"])
	clog.Info(globals.RootDir)
}
