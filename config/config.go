package config

import (
	"flag"
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/wuleying/silver-framework/globals"
)

type Config struct {
	Setting map[string]string
}

// 初始化配置
func ConfigInit() (Config, error) {
	var config Config

	flag.Parse()
	go_config, err := goconfig.LoadConfigFile(*globals.CONFIG_FILE_PATH)

	if err != nil {
		log.Printf("Read config file failed: %s", err)
		return config, err
	}

	log.Printf("Load config file success: %s", *globals.CONFIG_FILE_PATH)

	config.Setting, _ = go_config.GetSection("setting")

	return config, nil
}
