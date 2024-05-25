package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// setupLogger sets up the logger to write to a file in the logs directory
func setupLogger() (*os.File, error) {
	// Create the logs directory if it doesn't exist
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("could not create log directory: %v", err)
	}

	// Generate the log file name based on the current date and time
	logFileName := fmt.Sprintf("log_%s.txt", time.Now().Format("2006-01-02_15-04-05"))
	logFilePath := filepath.Join(logDir, logFileName)

	// Create or open the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("could not open log file: %v", err)
	}

	// Set the logger to write to the log file
	log.SetOutput(logFile)
	return logFile, nil
}
