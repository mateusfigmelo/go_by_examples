package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBasicFileOperations(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "test-basic-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory to temp dir
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Failed to restore working directory: %v", err)
		}
	}()

	// Run the function
	basicFileOperations()

	// Verify file contents
	data, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello, File I/O!\nThis is a test file.\n"
	if string(data) != expected {
		t.Errorf("Expected file contents %q, got %q", expected, string(data))
	}
}

func TestBufferedIO(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test-buffered-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Failed to restore working directory: %v", err)
		}
	}()

	// Run the function
	bufferedIO()

	// Verify file contents
	data, err := os.ReadFile("buffered.txt")
	if err != nil {
		t.Fatal(err)
	}

	expected := "Line 1: Buffered writing\nLine 2: More buffered content\n"
	if string(data) != expected {
		t.Errorf("Expected file contents %q, got %q", expected, string(data))
	}
}

func TestFileManipulation(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test-manipulation-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Failed to restore working directory: %v", err)
		}
	}()

	// Create initial test file
	if err := os.WriteFile("test.txt", []byte("test content"), 0644); err != nil {
		t.Fatal(err)
	}

	// Run the function
	fileManipulation()

	// Verify file was moved
	newPath := filepath.Join("testdir", "test.txt")
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		t.Error("File was not moved to testdir")
	}

	// Verify file contents preserved
	data, err := os.ReadFile(newPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "test content" {
		t.Errorf("File contents changed after move")
	}
}

func TestAdvancedFileOperations(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test-advanced-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Failed to restore working directory: %v", err)
		}
	}()

	// Run the function
	advancedFileOperations()

	// Note: We don't need to verify the temporary file contents
	// as it's automatically cleaned up, but we can verify the function
	// completes without error
}

func TestDirectoryOperations(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test-dir-ops-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(oldWd); err != nil {
			t.Logf("Failed to restore working directory: %v", err)
		}
	}()

	// Create test directory structure
	if err := os.MkdirAll("testdir", 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("testdir/test.txt", []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Run the function
	directoryOperations()

	// Verify directory contents
	entries, err := os.ReadDir("testdir")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in testdir, got %d", len(entries))
	}
	if entries[0].Name() != "test.txt" {
		t.Errorf("Expected file name 'test.txt', got %q", entries[0].Name())
	}
}
