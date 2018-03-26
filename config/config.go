package config

import (
	"flag"
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/wuleying/silver-framework/globals"
)

// Config struct
type Config struct {
	Setting map[string]string
}

// Init 初始化配置
func Init() (Config, error) {
	var config Config

	flag.Parse()
	configHandle, err := goconfig.LoadConfigFile(*globals.ConfigFilePath)

	if err != nil {
		log.Printf("Read config file failed: %s", err)
		return config, err
	}

	log.Printf("Load config file success: %s", *globals.ConfigFilePath)

	config.Setting, _ = configHandle.GetSection("setting")

	return config, nil
}
