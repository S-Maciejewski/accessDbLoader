package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
)

func InitLogFile(logFileName string) (file *os.File) {
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Error(err, fmt.Sprintf("Could not open / create logger file at %s", logFileName))
	}
	log.SetOutput(file)
	return
}

func CloseLogFile(logFile *os.File) {
	err := logFile.Close()
	if err != nil {
		Warning("Could not close logger file")
	}
}

func Error(err error, msg string) {
	log.Print(msg)
	fmt.Printf(errorColor, err.Error())
	fmt.Println()
	fmt.Printf(errorColor, msg)
	fmt.Println()
}

func ExecutionError(err error, statement string) {
	log.Printf("SQL execution error - %s - when executing statement %s\n", err.Error(), statement)
}

func Warning(msg string) {
	log.Println(msg)
	fmt.Printf(warningColor, msg)
	fmt.Println()
}

func Info(msg string) {
	log.Println(msg)
	fmt.Printf(infoColor, msg)
	fmt.Println()
}

func Notice(msg string) {
	fmt.Printf(noticeColor, msg)
	fmt.Println()
}
