package commands

import (
	"github.com/wuleying/silver-framework/version"
)

// PrintVersion 打印版本信息
func PrintVersion() error {
	println(version.Version)
	return nil
}
