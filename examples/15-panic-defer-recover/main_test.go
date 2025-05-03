package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
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

	// TODO: Add test assertions
	t.Logf("Program output: %s", output)
}

// Example test to demonstrate usage
func ExampleMain() {
	main()
	// TODO: Add expected output
	// Output:
}

func TestBasicPanic(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call basicPanic inside a recovery function
	didPanic := false
	var panicValue interface{}

	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
				panicValue = r
			}
		}()

		basicPanic()
	}()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}
	output := buf.String()

	// Check that we did indeed panic
	if !didPanic {
		t.Error("Expected basicPanic to panic, but it didn't")
	}

	// Check that the panic value is correct
	if panicValue != "something went wrong!" {
		t.Errorf("Expected panic value 'something went wrong!', got: %v", panicValue)
	}

	// Check that the deferred message was printed
	expectedOutput := "    Basic Panic Example:"
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain '%s', got: %s", expectedOutput, output)
	}

	// Check that the deferred message was printed
	deferredOutput := "        This will still run before the panic"
	if !strings.Contains(output, deferredOutput) {
		t.Errorf("Expected output to contain '%s', got: %s", deferredOutput, output)
	}

	// Check that the pre-panic message was printed
	prePanicOutput := "        About to panic..."
	if !strings.Contains(output, prePanicOutput) {
		t.Errorf("Expected output to contain '%s', got: %s", prePanicOutput, output)
	}
}
