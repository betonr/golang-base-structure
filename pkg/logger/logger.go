package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO]  ", log.LstdFlags|log.Lmsgprefix)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lmsgprefix)
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lmsgprefix)
}

func Info(msg string, args ...any) {
	infoLogger.Printf(msg, args...)
}

func Error(msg string, args ...any) {
	errorLogger.Printf(msg, args...)
}

func Debug(msg string, args ...any) {
	debugLogger.Printf(msg, args...)
}
