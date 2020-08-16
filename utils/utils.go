package utils

import (
    "os"
    "time"
    "strings"
)

func getHomePath() string {
    return os.Getenv("HOME")
}

func ExpandUser(path string) string {
    times := 1
    return strings.Replace(path, "~", getHomePath(), times)
}

func GetCurrentISOTimeString() string {
    currTime := time.Now()
    return currTime.Format("2006-01-02-150405")
}
