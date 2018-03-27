package config

import (
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/globals"
	"github.com/wuleying/silver-framework/utils"
)

// Config struct
type Config struct {
	Setting map[string]string
}

// Init 初始化配置
func Init() (Config, error) {
	var config Config
	configFilePath := globals.ConfigFileDefaultPath

	checkFile, err := utils.FileExists(globals.ConfigFilePath)

	if true == checkFile {
		configFilePath = globals.ConfigFilePath
	}

	configHandle, err := goconfig.LoadConfigFile(configFilePath)

	if err != nil {
		log.Printf("Read config file failed: %s", err)
		return config, err
	}

	clog.Info("Load config file success: %s", configFilePath)

	config.Setting, _ = configHandle.GetSection("setting")

	return config, nil
}
