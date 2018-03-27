package globals

// 全局常量
const (
	ProjectName = "silver-framework"
	Version     = "0.0.1"

	// 权限
	DBFileModel   = 0600
	FileReadMode  = 0644
	FileWriteMode = 0666
	DirReadMode   = 0755
	DirWriteMode  = 0777

	// Clog skip 级别
	ClogSkipDefault     = 0
	ClogSkipDisplayInfo = 2

	// 配置文件路径
	ConfigFilePath        = "config.ini"
	ConfigDefaultFilePath = "config.default.ini"
)
