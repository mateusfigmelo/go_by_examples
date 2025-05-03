// Package testutil provides shared testing utilities for go_by_examples.
package testutil

import (
	"bytes"
	"fmt"
	"os"
)

// CaptureOutput runs the given function and captures its stdout output.
// It returns the captured output as a string.
func CaptureOutput(f func()) string {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return fmt.Sprintf("ERROR CAPTURING OUTPUT: %v", err)
	}
	return buf.String()
}

// AssertOutput is a helper function to verify program output in tests.
func AssertOutput(got, want string) bool {
	return got == want
}

// ContainsOutput is a helper function to check if output contains expected string.
func ContainsOutput(output, substring string) bool {
	return bytes.Contains([]byte(output), []byte(substring))
}
