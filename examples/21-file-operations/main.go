package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// basicFileOperations demonstrates basic file reading and writing
func basicFileOperations() {
	fmt.Println("    Basic File Operations:")

	// Writing to a file
	content := []byte("Hello, File I/O!\nThis is a test file.\n")
	err := os.WriteFile("test.txt", content, 0644)
	if err != nil {
		fmt.Printf("        Error writing file: %v\n", err)
		return
	}
	fmt.Println("        File written successfully")

	// Reading from a file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("        Error reading file: %v\n", err)
		return
	}
	fmt.Printf("        File contents: %s", data)
}

// bufferedIO demonstrates buffered reading and writing
func bufferedIO() {
	fmt.Println("    Buffered I/O:")

	// Writing with a buffer
	file, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Printf("        Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString("Line 1: Buffered writing\n"); err != nil {
		fmt.Printf("        Error writing to file: %v\n", err)
		return
	}
	if _, err := writer.WriteString("Line 2: More buffered content\n"); err != nil {
		fmt.Printf("        Error writing to file: %v\n", err)
		return
	}
	if err := writer.Flush(); err != nil {
		fmt.Printf("        Error flushing buffer: %v\n", err)
		return
	}
	fmt.Println("        Buffered write complete")

	// Reading with a buffer
	file, err = os.Open("buffered.txt")
	if err != nil {
		fmt.Printf("        Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fmt.Println("        Reading line by line:")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("        Error reading line: %v\n", err)
			return
		}
		fmt.Printf("        %s", line)
	}
}

// fileManipulation demonstrates file manipulation operations
func fileManipulation() {
	fmt.Println("    File Manipulation:")

	// Create a directory
	err := os.MkdirAll("testdir", 0755)
	if err != nil {
		fmt.Printf("        Error creating directory: %v\n", err)
		return
	}
	fmt.Println("        Directory created")

	// Move a file
	err = os.Rename("test.txt", filepath.Join("testdir", "test.txt"))
	if err != nil {
		fmt.Printf("        Error moving file: %v\n", err)
		return
	}
	fmt.Println("        File moved to testdir")

	// Get file info
	info, err := os.Stat(filepath.Join("testdir", "test.txt"))
	if err != nil {
		fmt.Printf("        Error getting file info: %v\n", err)
		return
	}
	fmt.Printf("        File size: %d bytes\n", info.Size())
	fmt.Printf("        Modified: %v\n", info.ModTime())
	fmt.Printf("        Permissions: %v\n", info.Mode())
}

// advancedFileOperations demonstrates advanced file operations
func advancedFileOperations() {
	fmt.Println("    Advanced File Operations:")

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Printf("        Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	fmt.Printf("        Created temp file: %s\n", tempFile.Name())

	// Write some data
	data := []byte("This is temporary data\n")
	if _, err := tempFile.Write(data); err != nil {
		fmt.Printf("        Error writing to temp file: %v\n", err)
		return
	}

	// Seek to beginning of file
	if _, err := tempFile.Seek(0, 0); err != nil {
		fmt.Printf("        Error seeking in file: %v\n", err)
		return
	}

	// Read the data back
	readData := make([]byte, len(data))
	if _, err := tempFile.Read(readData); err != nil {
		fmt.Printf("        Error reading from temp file: %v\n", err)
		return
	}
	fmt.Printf("        Read from temp file: %s", string(readData))
}

// directoryOperations demonstrates directory operations
func directoryOperations() {
	fmt.Println("    Directory Operations:")

	// List directory contents
	entries, err := os.ReadDir("testdir")
	if err != nil {
		fmt.Printf("        Error reading directory: %v\n", err)
		return
	}

	fmt.Println("        Directory contents:")
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("        Error getting entry info: %v\n", err)
			continue
		}
		fmt.Printf("        - %s (%d bytes)\n", entry.Name(), info.Size())
	}

	// Walk directory tree
	fmt.Println("        Walking directory tree:")
	err = filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("        - %s\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("        Error walking directory: %v\n", err)
	}
}

// cleanup removes test files and directories
func cleanup() {
	fmt.Println("    Cleaning up...")
	os.RemoveAll("testdir")
	os.Remove("buffered.txt")
	fmt.Println("        Cleanup complete")
}

func main() {
	fmt.Println("=== Go File Operations Examples ===")

	fmt.Println("\n1. Basic File Operations:")
	basicFileOperations()

	fmt.Println("\n2. Buffered I/O:")
	bufferedIO()

	fmt.Println("\n3. File Manipulation:")
	fileManipulation()

	fmt.Println("\n4. Advanced File Operations:")
	advancedFileOperations()

	fmt.Println("\n5. Directory Operations:")
	directoryOperations()

	fmt.Println("\n6. Cleanup:")
	cleanup()
}
