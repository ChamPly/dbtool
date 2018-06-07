// +build linux
package log

import (
	"fmt"

	"github.com/ChamPly/dbtool/utility"
)

func Infof(format string, params ...interface{}) {
	fmt.Printf("%s "+format+"\n", utility.GetFormatTime(), params)
	return
}

func Info(content interface{}) {
	fmt.Printf("%s %v", utility.GetFormatTime(), content)
	return
}

func Errorf(format string, params ...interface{}) {
	fmt.Printf("%s "+format+"\n", utility.GetFormatTime(), params)
	return
}

func Error(content interface{}) {
	fmt.Printf("%s %v", utility.GetFormatTime(), content)
	return
}
