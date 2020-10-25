package datetime

import (
	"time"
)

func GetCurrentISOTimeString() string {
	currTime := time.Now()
	return currTime.Format("2006-01-02-150405")
}