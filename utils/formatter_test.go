package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReadableFloat(t *testing.T) {
	Convey("格式化浮点数", t, func() {
		So(ReadableFloat(1.001), ShouldEqual, "1.001")
		So(ReadableFloat(100.01), ShouldEqual, "100.01")
		So(ReadableFloat(1000.99), ShouldEqual, "1000.99")
	})
}
