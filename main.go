package main

import (
	"fmt"
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

	cfg, err := config.Init()
	exceptions.CheckError(err)

	http := admin.HTTP{Config: &cfg}
	http.Init()
}
