package util

import "time"

func UnixMillTime(oldUnix int64) int64{
	return oldUnix/int64(time.Millisecond)
}
