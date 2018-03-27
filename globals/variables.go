package globals

import (
	"github.com/wuleying/silver-framework/utils"
	"time"
)

// 全局变量
var (
	// 根目录
	RootDir = utils.FileGetParentDirectory(utils.FileGetCurrentDirectory())
	// 当前时间
	CurrentTime = time.Now().String()
)
