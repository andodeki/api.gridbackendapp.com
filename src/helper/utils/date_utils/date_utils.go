package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow return time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString func that get the dates
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat func that get the dates
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
