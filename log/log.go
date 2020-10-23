package log

import "fmt"

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
)

func Error(err error, msg string) {
	fmt.Printf(errorColor, err.Error())
	fmt.Println()
	fmt.Printf(errorColor, msg)
	fmt.Println()
}

func WarningMessage(msg string) {
	fmt.Printf(warningColor, msg)
	fmt.Println()
}
