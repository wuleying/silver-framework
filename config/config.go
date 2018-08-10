package config

import (
	"github.com/Unknwon/goconfig"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/globals"
	"github.com/wuleying/silver-framework/utils"
)

// Config struct
type Config struct {
	Setting map[string]string
	Metrics map[string]string
}

// Init 初始化配置
func Init() (Config, error) {
	var config Config
	configFilePath := globals.ConfigDefaultFilePath

	checkFile, err := utils.FileExists(globals.ConfigFilePath)
	exceptions.CheckError(err)

	if checkFile {
		configFilePath = globals.ConfigFilePath
	}

	configHandle, err := goconfig.LoadConfigFile(configFilePath)
	exceptions.CheckError(err)

	clog.Info("Load config file success: %s", configFilePath)

	config.Setting, _ = configHandle.GetSection("setting")
	config.Metrics, _ = configHandle.GetSection("metrics")

	return config, nil
}
