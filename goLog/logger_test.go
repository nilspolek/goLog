package goLog

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewStdLogger(t *testing.T) {
	logger, err := NewStdLogger()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if logger.toFile != false {
		t.Errorf("expected toFile to be false, got %v", logger.toFile)
	}

	if logger.toStdout != true {
		t.Errorf("expected toStdout to be true, got %v", logger.toStdout)
	}
}

func TestNewLogger(t *testing.T) {
	// Create a temporary file for testing
	file, err := ioutil.TempFile("", "test_log")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up the file after test
	defer file.Close()

	logger, err := NewLogger(file.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if logger.toFile != true {
		t.Errorf("expected toFile to be true, got %v", logger.toFile)
	}

	if logger.toStdout != false {
		t.Errorf("expected toStdout to be false, got %v", logger.toStdout)
	}

	logger.Info("Test Info Message")
	logger.Close()

	// Check if log message was written to the file
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "Test Info Message") {
		t.Errorf("expected log to contain 'Test Info Message', but got %s", string(content))
	}
}

func TestLogLevels(t *testing.T) {
	// Create a temporary file for testing
	file, err := ioutil.TempFile("", "test_log_levels")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	logger, err := NewLogger(file.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer logger.Close()

	logger.Debug("Test Debug Message")
	logger.Info("Test Info Message")
	logger.Warning("Test Warning Message")
	logger.Error("Test Error Message")

	// Read the file content to verify logs
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	logContent := string(content)
	if !strings.Contains(logContent, "Test Debug Message") {
		t.Errorf("expected log to contain 'Test Debug Message', but got %s", logContent)
	}
	if !strings.Contains(logContent, "Test Info Message") {
		t.Errorf("expected log to contain 'Test Info Message', but got %s", logContent)
	}
	if !strings.Contains(logContent, "Test Warning Message") {
		t.Errorf("expected log to contain 'Test Warning Message', but got %s", logContent)
	}
	if !strings.Contains(logContent, "Test Error Message") {
		t.Errorf("expected log to contain 'Test Error Message', but got %s", logContent)
	}
}

func TestCloseLogger(t *testing.T) {
	// Create a temporary file for testing
	file, err := ioutil.TempFile("", "test_close_logger")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	logger, err := NewLogger(file.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = logger.Close()
	if err != nil {
		t.Errorf("expected no error on close, got %v", err)
	}

	// Try closing again to make sure it doesn't return an error
	err = logger.Close()
	if err != nil {
		t.Errorf("expected no error on double close, got %v", err)
	}
}

func TestVisual(t *testing.T) {
	logger, _ := NewStdLogger()
	logger.Info("Test Info Message")
	logger.Warning("Test Warning Message")
	logger.Error("Test Error Message")
	logger.Debug("Test Debug Message")
	err := logger.Close()
	if err != nil {
		t.Errorf("expected no error on close, got %v", err)
	}
}
