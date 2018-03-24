package globals

import (
	"flag"
	"github.com/wuleying/silver-framework/utils"
	"time"
)

// 全局变量
var (
	// 根目录
	ROOT_DIR = utils.FileGetParentDirectory(utils.FileGetCurrentDirectory())
	// 当前时间
	CURRENT_TIME = time.Now().String()
	// 配置文件路径
	CONFIG_FILE_PATH = flag.String("config", "config-dev.ini", "config file path")
)
