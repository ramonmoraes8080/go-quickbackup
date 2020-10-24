package output

import (
	"fmt"
)

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