package util

import (
	"strconv"
	"strings"
)

func Int64ToString (a int64) string{
	return strconv.FormatInt(a,10)
}


func CleanNewlineChart (in string) string {
	return strings.Replace(
		strings.Replace(in,"\n","|^---^|",-1),
		"\r","|^---^|",-1)
}
