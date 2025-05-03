package main

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed data/*
var embeddedFiles embed.FS

// lineFilter demonstrates processing input line by line
func lineFilter() {
	fmt.Println("    Line Filter Example:")

	// Create a test file with some lines
	testLines := []string{
		"First line",
		"SECOND LINE",
		"third line",
		"Fourth Line",
		"FIFTH LINE",
	}

	err := os.WriteFile("input.txt", []byte(strings.Join(testLines, "\n")), 0644)
	if err != nil {
		fmt.Printf("        Error creating test file: %v\n", err)
		return
	}

	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("        Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create scanner for line-by-line reading
	scanner := bufio.NewScanner(file)
	fmt.Println("        Processing lines:")

	// Process each line
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		// Example transformations
		if strings.ToUpper(line) == line {
			fmt.Printf("        %d: [UPPERCASE] %s\n", lineNum, line)
		} else if strings.ToLower(line) == line {
			fmt.Printf("        %d: [lowercase] %s\n", lineNum, line)
		} else {
			fmt.Printf("        %d: [Mixed Case] %s\n", lineNum, line)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("        Error reading file: %v\n", err)
	}
}

// filePathOperations demonstrates working with file paths
func filePathOperations() {
	fmt.Println("    File Path Operations:")

	// Join paths
	path := filepath.Join("dir", "subdir", "file.txt")
	fmt.Printf("        Joined path: %s\n", path)

	// Split path
	dir, file := filepath.Split(path)
	fmt.Printf("        Split path - Dir: %s, File: %s\n", dir, file)

	// Get absolute path
	absPath, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("        Error getting absolute path: %v\n", err)
	} else {
		fmt.Printf("        Absolute path: %s\n", absPath)
	}

	// Clean path
	messyPath := "dir/../dir/./subdir//file.txt"
	cleanPath := filepath.Clean(messyPath)
	fmt.Printf("        Cleaned path: %s\n", cleanPath)

	// File extension
	ext := filepath.Ext("document.pdf")
	fmt.Printf("        File extension: %s\n", ext)

	// Base name
	base := filepath.Base(path)
	fmt.Printf("        Base name: %s\n", base)
}

// directoryWalking demonstrates advanced directory operations
func directoryWalking() {
	fmt.Println("    Directory Walking:")

	// Create test directory structure
	dirs := []string{
		"testdir/a/b/c",
		"testdir/x/y/z",
		"testdir/1/2/3",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("        Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create some test files
	files := map[string]string{
		"testdir/a/file1.txt": "Content 1",
		"testdir/x/file2.txt": "Content 2",
		"testdir/1/file3.txt": "Content 3",
	}

	for path, content := range files {
		err := os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("        Error creating file %s: %v\n", path, err)
			return
		}
	}

	// Walk the directory tree
	fmt.Println("        Walking directory tree:")
	err := filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Printf("        DIR: %s\n", path)
		} else {
			fmt.Printf("        FILE: %s (%d bytes)\n", path, info.Size())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("        Error walking directory: %v\n", err)
	}
}

// tempFilesAndDirs demonstrates working with temporary files and directories
func tempFilesAndDirs() {
	fmt.Println("    Temporary Files and Directories:")

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "example-*")
	if err != nil {
		fmt.Printf("        Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)
	fmt.Printf("        Created temp directory: %s\n", tempDir)

	// Create temporary file in the directory
	tempFile, err := os.CreateTemp(tempDir, "tempfile-*.txt")
	if err != nil {
		fmt.Printf("        Error creating temp file: %v\n", err)
		return
	}
	defer tempFile.Close()
	fmt.Printf("        Created temp file: %s\n", tempFile.Name())

	// Write some data
	data := []byte("This is temporary data\n")
	if _, err := tempFile.Write(data); err != nil {
		fmt.Printf("        Error writing to temp file: %v\n", err)
		return
	}

	// List contents of temp directory
	files, err := os.ReadDir(tempDir)
	if err != nil {
		fmt.Printf("        Error reading temp directory: %v\n", err)
		return
	}

	fmt.Println("        Temporary directory contents:")
	for _, file := range files {
		fmt.Printf("        - %s\n", file.Name())
	}
}

// embedDirective demonstrates using the embed directive
func embedDirective() {
	fmt.Println("    Embed Directive Example:")

	// Read embedded config file
	configData, err := embeddedFiles.ReadFile("data/config.txt")
	if err != nil {
		fmt.Printf("        Error reading embedded config: %v\n", err)
		return
	}
	fmt.Println("        Embedded config file contents:")
	fmt.Printf("%s\n", configData)

	// Read and parse embedded template
	tmpl, err := template.ParseFS(embeddedFiles, "data/template.html")
	if err != nil {
		fmt.Printf("        Error parsing embedded template: %v\n", err)
		return
	}

	// Use the template
	data := struct {
		Title  string
		Header string
		Items  []string
	}{
		Title:  "Embedded Template Demo",
		Header: "File Utilities Example",
		Items:  []string{"Item 1", "Item 2", "Item 3"},
	}

	fmt.Println("        Rendered template:")
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		fmt.Printf("        Error executing template: %v\n", err)
	}
}

// cleanup removes test files and directories
func cleanup() {
	fmt.Println("    Cleaning up...")
	os.Remove("input.txt")
	os.RemoveAll("testdir")
	fmt.Println("        Cleanup complete")
}

func main() {
	fmt.Println("=== Go File Utilities Examples ===")

	fmt.Println("\n1. Line Filter:")
	lineFilter()

	fmt.Println("\n2. File Path Operations:")
	filePathOperations()

	fmt.Println("\n3. Directory Walking:")
	directoryWalking()

	fmt.Println("\n4. Temporary Files and Directories:")
	tempFilesAndDirs()

	fmt.Println("\n5. Embed Directive:")
	embedDirective()

	fmt.Println("\n6. Cleanup:")
	cleanup()
}
