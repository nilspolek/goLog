package goLog

import (
	"fmt"
	"os"
)

// ExampleNewStdLogger demonstrates how to use the NewStdLogger function.
func ExampleNewStdLogger() {
	// Create a standard logger (logs to stdout)
	logger, err := NewStdLogger()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Use the logger to log messages
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	// Output:
	// [INFO] This is an info message
	// [WARNING] This is a warning message
	// [ERROR] This is an error message
}

// ExampleNewLogger demonstrates how to use the NewLogger function.
func ExampleNewLogger() {
	// Create a logger that writes to a file
	filePath := "example.log"
	logger, err := NewLogger(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer os.Remove(filePath) // Clean up file after example
	defer logger.Close()

	// Log some messages
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")

	// Output to console will not occur as the logger is set to file-only.
	// To verify, check the contents of example.log.
}
