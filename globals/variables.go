package globals

import (
	"flag"
	"github.com/wuleying/silver-framework/utils"
	"time"
)

// 全局变量
var (
	// 根目录
	RootDir = utils.FileGetParentDirectory(utils.FileGetCurrentDirectory())
	// 当前时间
	CurrentTime = time.Now().String()
	// 配置文件路径
	ConfigFilePath = flag.String("config", "config-dev.ini", "config file path")
)
