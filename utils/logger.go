package utils

import "log"

type Logger struct {
	moduleName string
}

func NewLogger(moduleName string) *Logger {
	return &Logger{moduleName: "[" + moduleName + "]"}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(l.moduleName+"(fatal) "+format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Printf(l.moduleName+"(error) "+format, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	log.Printf(l.moduleName+"(info) "+format, v...)
}

var logger *Logger

func init() {
	logger = NewLogger("utils")
}
