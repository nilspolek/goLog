package goLog

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestLogOnError(t *testing.T) {
	if LogOnError("Hallo", errors.New("Test")) != "Hallo" {
		t.Fatalf("LogOnError() failed.")
	}
}

func TestLogToFile(t *testing.T) {
	file, err := os.CreateTemp("", "log_test")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	ToFile(file)

	Log(INFO, "Test file logging")

	file.Close()

	data, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("could not read temp file: %v", err)
	}

	expected := "[INFO] Test file logging"
	if !bytes.Contains(data, []byte(expected)) {
		t.Errorf("Log to file = %v, want %v", string(data), expected)
	}
}

func TestClose(t *testing.T) {
	file, err := os.CreateTemp("", "log_test")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	ToFile(file)

	if err := file.Close(); err != nil {
		t.Fatalf("could not close temp file: %v", err)
	}
}

func TestLogLevelToString(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected string
	}{
		{DEBUG, "DEBUG"},
		{INFO, "INFO"},
		{WARNING, "WARNING"},
		{ERROR, "ERROR"},
	}

	for _, tt := range tests {
		t.Run(levelToString(tt.level), func(t *testing.T) {
			if got := levelToString(tt.level); got != tt.expected {
				t.Errorf("levelToString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
