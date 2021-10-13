package logger

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func NewLogger() Logger {
	return Logger{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Llongfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC),
	}
}
