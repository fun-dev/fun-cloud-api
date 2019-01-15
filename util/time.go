package util

import "time"

const (
	format = "20060102150405"
)

func GetNowString() string {
	return time.Now().Format(format)
}
