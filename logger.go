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

	// LOW Loglevel means only Error and WARNING get Logged
	LOW
	// MEDIUM Loglevel means Info, WARNING and Error get Logged
	MEDIUM
	// HIGH Loglevel means everything gets Logged
	HIGH

	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
)

var (
	file   *os.File
	toFile = false
	// IsDebug Changes if debug messages are displayed.
	LoggingLevel LogLevel
)

func ToFile(file2 *os.File) {
	toFile = true
	file = file2
}

func Logf(level LogLevel, msg string, args []any) {
	Log(level, fmt.Sprintf(msg, args...))
}

func Log(level LogLevel, msg string) {

	color := getColor(level)
	resetColor := colorReset

	logMsg := fmt.Sprintf(" [%s] %s", levelToString(level), msg)

	log.Printf("%s%s%s", color, logMsg, resetColor)

	if toFile && file != nil {
		file.WriteString(logMsg)
	}
}

func setLogLevel(level LogLevel) {
	LoggingLevel = level
}

func Debug(msg string, args ...any) {
	if LoggingLevel == HIGH {
		return
	}
	Logf(DEBUG, msg, args)
}

func Info(msg string, args ...any) {
	if LoggingLevel == HIGH || LoggingLevel == MEDIUM {
		return
	}
	Logf(INFO, msg, args)
}

func Warning(msg string, args ...any) {
	Logf(WARNING, msg, args)
}

func Error(msg string, args ...any) {
	Logf(ERROR, msg, args)
}

func LogOnError[T any](a T, err error) T {
	if err != nil {
		Error(err.Error())
	}
	return a
}

func ExitOnError[T any](a T, err error) T {
	if err != nil {
		Error(err.Error())
		os.Exit(1)
	}
	return a
}

func levelToString(level LogLevel) string {
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

func getColor(level LogLevel) string {
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

func Close() error {
	if toFile && file != nil {
		err := file.Close()
		file = nil
		return err
	}
	return nil
}
