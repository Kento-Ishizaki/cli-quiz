package main

import (
	"log"
	"os"
)

// LogError logs the error and exits the program
func LogError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
		os.Exit(1)
	}
}
