package utils

import "time"

// return the millisecond timestamp
func GetNowTimestamp() int64 {
	return time.Now().UTC().UnixMilli()
}
