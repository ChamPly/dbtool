package utility

import (
	"strconv"
	"strings"
	"time"
)

// StringToInt string to int, if string is empty, return 0, else return result and error
func StringToInt(str string) (result int, err error) {
	if strings.EqualFold(str, "") {
		return
	}

	result, err = strconv.Atoi(str)
	return
}

// GetFormatTime return format time result
func GetFormatTime() (t string) {
	return time.Now().Format("2006-01-02 15:04:05")
}
