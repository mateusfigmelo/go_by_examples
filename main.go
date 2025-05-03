package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Color codes for terminal output
const (
	colorRed    = "\033[0;31m"
	colorGreen  = "\033[0;32m"
	colorYellow = "\033[0;33m"
	colorBlue   = "\033[0;34m"
	colorReset  = "\033[0m"
)

// Global options
var (
	verbose  bool
	coverage bool
	example  string
)

// Paths
var (
	projectRoot string
	coverageDir string
)

func init() {
	// Parse command line flags
	flag.BoolVar(&verbose, "v", false, "Run tests with verbose output")
	flag.BoolVar(&verbose, "verbose", false, "Run tests with verbose output")
	flag.BoolVar(&coverage, "c", false, "Generate test coverage reports")
	flag.BoolVar(&coverage, "coverage", false, "Generate test coverage reports")

	// Custom usage message
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] [EXAMPLE_ID]\n", os.Args[0])
		fmt.Println("Run tests for Go by Examples")
		fmt.Println("")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Printf("  %s                Run all tests\n", os.Args[0])
		fmt.Printf("  %s 01-hello-world Run tests for the hello-world example\n", os.Args[0])
		fmt.Printf("  %s -c 24-cli      Run tests for the CLI example with coverage\n", os.Args[0])
		os.Exit(1)
	}

	// Find project root
	var err error
	projectRoot, err = os.Getwd()
	if err != nil {
		logError("Failed to get current directory: %v", err)
		os.Exit(1)
	}

	// Set coverage directory
	coverageDir = filepath.Join(projectRoot, "coverage")
}

// log prints a timestamped message to the console
func log(format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] %s\n", timestamp, message)
}

// logError prints an error message in red
func logError(format string, args ...interface{}) {
	log(colorRed+format+colorReset, args...)
}

// logSuccess prints a success message in green
func logSuccess(format string, args ...interface{}) {
	log(colorGreen+format+colorReset, args...)
}

// logWarning prints a warning message in yellow
func logWarning(format string, args ...interface{}) {
	log(colorYellow+format+colorReset, args...)
}

// logInfo prints an info message in blue
func logInfo(format string, args ...interface{}) {
	log(colorBlue+format+colorReset, args...)
}

// checkGoVersion verifies Go is installed
func checkGoVersion() {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		logError("Error: Go is not installed or not in the PATH")
		os.Exit(1)
	}
	logInfo("Using Go version: %s", strings.TrimSpace(string(output)))
}

// prepareCoverage creates the directory for coverage reports
func prepareCoverage() {
	if coverage {
		err := os.MkdirAll(coverageDir, 0755)
		if err != nil {
			logError("Failed to create coverage directory: %v", err)
			os.Exit(1)
		}
		logInfo("Coverage reports will be saved to: %s", coverageDir)
	}
}

// runExampleTests runs tests for a specific example
func runExampleTests(exampleID string) (bool, error) {
	dir := filepath.Join("examples", exampleID)

	// Skip if directory doesn't exist
	if _, err := os.Stat(filepath.Join(projectRoot, dir)); os.IsNotExist(err) {
		logError("Error: Directory %s does not exist", dir)
		return false, err
	}

	// Skip if main.go doesn't exist
	mainFile := filepath.Join(projectRoot, dir, "main.go")
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		logWarning("Warning: %s does not contain main.go, skipping", dir)
		return true, nil
	}

	logSuccess("Running tests for %s...", exampleID)

	// Build test command
	args := []string{"test"}
	if verbose {
		args = append(args, "-v")
	}

	if coverage {
		coverageFile := filepath.Join(coverageDir, exampleID+".out")
		args = append(args, "-coverprofile="+coverageFile, "-covermode=atomic")
	}

	args = append(args, "./...")

	// Change to example directory and run tests
	cmd := exec.Command("go", args...)
	cmd.Dir = filepath.Join(projectRoot, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	// Generate HTML coverage report if requested and tests passed
	if coverage && err == nil {
		coverageFile := filepath.Join(coverageDir, exampleID+".out")
		htmlFile := filepath.Join(coverageDir, exampleID+".html")

		coverCmd := exec.Command("go", "tool", "cover", "-html="+coverageFile, "-o="+htmlFile)
		coverErr := coverCmd.Run()
		if coverErr == nil {
			logInfo("Coverage report generated: %s", htmlFile)
		} else {
			logWarning("Failed to generate coverage report: %v", coverErr)
		}
	}

	// Report test results
	if err == nil {
		logSuccess("✓ Tests passed for %s", exampleID)
		return true, nil
	} else {
		logError("✗ Tests failed for %s", exampleID)
		return false, err
	}
}

// runAllTests runs tests for all examples
func runAllTests() bool {
	logSuccess("Running all tests...")
	allPassed := true
	var failedExamples []string

	// Find all example directories
	examplesDir := filepath.Join(projectRoot, "examples")
	entries, err := os.ReadDir(examplesDir)
	if err != nil {
		logError("Failed to read examples directory: %v", err)
		return false
	}

	// Process each example directory
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		// Check if it follows the pattern NN-name
		if !strings.HasPrefix(name, "0") && !strings.HasPrefix(name, "1") && !strings.HasPrefix(name, "2") && !strings.HasPrefix(name, "3") {
			continue
		}

		// Skip if no main.go file
		mainFile := filepath.Join(examplesDir, name, "main.go")
		if _, err := os.Stat(mainFile); os.IsNotExist(err) {
			logWarning("Skipping %s: no main.go found", name)
			continue
		}

		// Run tests for this example
		passed, _ := runExampleTests(name)
		if !passed {
			allPassed = false
			failedExamples = append(failedExamples, name)
		}
	}

	// Print summary
	fmt.Println()
	if allPassed {
		logSuccess("All tests passed successfully!")
	} else {
		logError("The following examples had failing tests:")
		for _, example := range failedExamples {
			fmt.Printf("%s  - %s%s\n", colorRed, example, colorReset)
		}
	}

	return allPassed
}

func main() {
	// Parse flags
	flag.Parse()

	// Check for example ID as a non-flag argument
	if flag.NArg() > 0 {
		example = flag.Arg(0)
	}

	// Check Go installation
	checkGoVersion()

	// Prepare coverage dir if needed
	prepareCoverage()

	var success bool

	// Run tests for specific example or all examples
	if example != "" {
		success, _ = runExampleTests(example)
	} else {
		success = runAllTests()
	}

	// Exit with appropriate status code
	if !success {
		os.Exit(1)
	}
}
