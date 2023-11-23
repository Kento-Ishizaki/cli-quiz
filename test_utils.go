package main

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLogError(t *testing.T) {
	// Create a fake error
	err := errors.New("this is a test error")

	// Redirect the standard output to a buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function with the fake error
	LogError(err, "Test message")

	// Restore the standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the buffer contents
	out, _ := io.ReadAll(r)
	// Check if the error message is in the buffer
	expected := "Test message: This is a test error"
	if !strings.Contains(string(out), expected) {
		t.Errorf("LogError() = %s; want %s", string(out), expected)
	}
}
