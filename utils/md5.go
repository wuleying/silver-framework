package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 函数
func Md5(raw string) string {
	h := md5.Sum([]byte(raw))
	return hex.EncodeToString(h[:])
}
