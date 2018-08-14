package config

import (
	"github.com/Unknwon/goconfig"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/globals"
	"github.com/wuleying/silver-framework/utils"
)

type Config map[string]map[string]string

// Init 初始化配置
func Init() (Config, error) {
	configs := make(Config)
	configFilePath := globals.ConfigDefaultFilePath

	checkFile, err := utils.FileExists(globals.ConfigFilePath)
	exceptions.CheckError(err)

	if checkFile {
		configFilePath = globals.ConfigFilePath
	}

	configHandle, err := goconfig.LoadConfigFile(configFilePath)
	exceptions.CheckError(err)

	clog.Info("Load config file success: %s", configFilePath)

	sectionList := configHandle.GetSectionList()

	for i := 0; i < len(sectionList); i++ {
		configs[sectionList[i]], _ = configHandle.GetSection(sectionList[i])
	}

	return configs, nil
}
