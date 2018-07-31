package globals

// 全局常量
const (
	// 系统版本
	Version = "0.0.1"

	// 配置文件路径
	ConfigFilePath        = "config.ini"
	ConfigDefaultFilePath = "config.default.ini"

	// Clog skip 级别
	ClogSkipDefault     = 0
	ClogSkipDisplayInfo = 2
)

// 权限
const (
	DatabaseFileModel = 0600
	FileReadMode      = 0644
	FileWriteMode     = 0666
	DirReadMode       = 0755
	DirWriteMode      = 0777
)

// 存储大小
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)
