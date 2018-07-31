package utils

import "github.com/petermattis/goid"

// GetGoID 获取 goroutine id
func GetGoID() int64 {
	return goid.Get()
}
