package helper

import (
	"fmt"
	"gitee.com/up-zero/up-admin/define"
	"runtime"
	"strconv"
	"time"
)

func dataFormat(showDetail bool, format string, v ...interface{}) string {
	prefix := "[" + define.FrameName + "] " + time.Now().Format(define.DateTimeLayout) + " "
	if showDetail {
		_, file, line, _ := runtime.Caller(2)
		prefix += "file: " + file + " line: " + strconv.Itoa(line) + " ==> "
	}
	return prefix + fmt.Sprintf(format, v...)
}

// Info INFO
func Info(format string, v ...interface{}) {
	fmt.Printf("\033[32m%s\033[0m\n", dataFormat(false, format, v...))
}

// Error ERROR
func Error(format string, v ...interface{}) {
	fmt.Printf("\033[31m%s\033[0m\n", dataFormat(true, format, v...))
}
