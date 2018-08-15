package utils

import (
	"crypto/md5"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"testing"
)

func origMd5(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TestMd5(t *testing.T) {
	Convey("md5字符串", t, func() {
		So(Md5("luo"), ShouldEqual, origMd5("luo"))
		So(Md5("Luo"), ShouldNotEqual, origMd5("luo"))
	})
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5("luo")
	}
}

func BenchmarkOrigMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		origMd5("luo")
	}
}
