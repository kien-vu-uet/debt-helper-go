package bootstrap

import (
	"log"
	"os"
)

// Logger defines a simple logger interface
type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewLogger creates a new Logger instance
func NewLogger() *Logger {
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}

// Info logs informational messages
func (l *Logger) Info(message string) {
	l.InfoLogger.Println(message)
}

// Error logs error messages
func (l *Logger) Error(message string, err error) {
	if err != nil {
		l.ErrorLogger.Printf("%s: %v", message, err)
	} else {
		l.ErrorLogger.Println(message)
	}
}
