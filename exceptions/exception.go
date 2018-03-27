package exceptions

import (
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/globals"
	"log"
	"runtime"
)

// CheckError 错误检查
func CheckError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Println(file)
		log.Println(line)

		clog.Fatal(globals.ClogSkipDefault, "%s:%d %s", file, line, err.Error())
	}
}
