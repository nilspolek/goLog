package goLog

import (
	"fmt"
	"log"
	"os"
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
)

type Logger struct {
	file     *os.File
	toFile   bool
	toStdout bool
}

func NewStdLogger() (*Logger, error) {
	return &Logger{
		file:     nil,
		toFile:   false,
		toStdout: true,
	}, nil
}

func NewLogger(filePath string) (*Logger, error) {
	var file *os.File
	var err error

	if filePath != "" {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("could not open log file: %v", err)
		}
	}
	return &Logger{
		file:     file,
		toFile:   filePath != "",
		toStdout: false,
	}, nil
}

func (l *Logger) Logf(level LogLevel, msg string, args []any) {
	l.Log(level, fmt.Sprintf(msg, args...))
}

func (l *Logger) Log(level LogLevel, msg string) {

	color := l.getColor(level)
	resetColor := colorReset

	logMsg := fmt.Sprintf(" [%s] %s", l.levelToString(level), msg)

	if l.toStdout {
		log.Printf("%s%s%s", color, logMsg, resetColor)
	}

	if l.toFile && l.file != nil {
		_, _ = l.file.WriteString(logMsg + "\n")
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.Logf(DEBUG, msg, args)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logf(INFO, msg, args)
}

func (l *Logger) Warning(msg string, args ...any) {
	l.Logf(WARNING, msg, args)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logf(ERROR, msg, args)
}

func (l *Logger) levelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) getColor(level LogLevel) string {
	switch level {
	case DEBUG:
		return colorBlue
	case INFO:
		return colorGreen
	case WARNING:
		return colorYellow
	case ERROR:
		return colorRed
	default:
		return colorReset
	}
}

func (l *Logger) Close() error {
	if l.toFile && l.file != nil {
		err := l.file.Close()
		l.file = nil
		return err
	}
	return nil
}
