package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestStringSub
func TestStringSub(t *testing.T) {
	Convey("截取字符串", t, func() {
		So(StringSub("hello world", 0, 3), ShouldEqual, "hel")
		So(StringSub("你好世界", 0, 3), ShouldEqual, "你好世")
	})
}
