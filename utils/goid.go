package utils

import "github.com/petermattis/goid"

// GetGoId 获取 goroutine id
func GetGoId() int64 {
	return goid.Get()
}
