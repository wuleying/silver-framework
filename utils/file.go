package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// FileGetCurrentDirectory 获取当前目录
func FileGetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Get current directory failed.", err.Error())
	}

	return strings.Replace(dir, "\\", "/", -1)
}

// FileGetParentDirectory 获取上级目录
func FileGetParentDirectory(directory string) string {
	return StringSub(directory, 0, strings.LastIndex(directory, "/"))
}

// FileGetName 获取文件名称
func FileGetName(filePath string) string {
	return path.Base(filePath)
}

// FileGetSize 获取取文件大小
func FileGetSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	fileSize := fileInfo.Size()
	return fileSize, nil
}
