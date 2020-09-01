/*
Copyright © 2020 Ramon Moraes <ramonmoraes8080@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// https://gist.github.com/ik5/d8ecde700972d4378d87
const (
	SuccessColor = "\033[1;32m%s\033[0m"
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	MagentaColor = "\033[1;35m%s\033[0m"
)

func LoggerSuccess(msg string) {
	fmt.Println(fmt.Sprintf(SuccessColor, msg))
}

func LoggerInfo(msg string) {
	fmt.Println(fmt.Sprintf(InfoColor, msg))
}

func LoggerWarning(msg string) {
	fmt.Println(fmt.Sprintf(WarningColor, msg))
}

func LoggerError(msg string) {
	fmt.Println(fmt.Sprintf(ErrorColor, msg))
}

func LoggerMagenta(msg string) {
	fmt.Println(fmt.Sprintf(MagentaColor, msg))
}

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
