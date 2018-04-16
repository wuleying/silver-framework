package commands

import "github.com/wuleying/silver-framework/globals"

// PrintVersion 打印版本信息
func PrintVersion() error {
	println(globals.Version)
	return nil
}
