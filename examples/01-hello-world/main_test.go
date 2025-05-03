package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the main program
	main()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}
	output := strings.TrimSpace(buf.String())

	// Verify output
	expected := "Hello, World!"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}

// Example test to demonstrate the program's usage
func ExampleMain() {
	main()
	// Output: Hello, World!
}
