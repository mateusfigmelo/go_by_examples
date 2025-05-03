package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runTests() {
	fmt.Println("\nRunning Tests:")
	cmd := exec.Command("go", "test", "./calculator", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running tests: %v\n", err)
	}
}

func runBenchmarks() {
	fmt.Println("\nRunning Benchmarks:")
	cmd := exec.Command("go", "test", "./calculator", "-bench=.", "-benchmem")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running benchmarks: %v\n", err)
	}
}

func runCoverage() {
	fmt.Println("\nRunning Coverage Analysis:")

	// Generate coverage profile
	cmd := exec.Command("go", "test", "./calculator", "-coverprofile=coverage.out")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error generating coverage profile: %v\n", err)
		return
	}

	// Display coverage report
	cmd = exec.Command("go", "tool", "cover", "-func=coverage.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error displaying coverage report: %v\n", err)
	}

	// Clean up coverage file
	os.Remove("coverage.out")
}

func main() {
	fmt.Println("=== Go Testing and Benchmarking Examples ===")

	fmt.Println("\nThis example demonstrates:")
	fmt.Println("1. Unit Testing")
	fmt.Println("   - Basic tests")
	fmt.Println("   - Table-driven tests")
	fmt.Println("   - Subtests")
	fmt.Println("   - Error testing")

	fmt.Println("\n2. Benchmarking")
	fmt.Println("   - Performance measurements")
	fmt.Println("   - Memory allocations")
	fmt.Println("   - Comparative benchmarks")

	fmt.Println("\n3. Coverage Analysis")
	fmt.Println("   - Code coverage reporting")
	fmt.Println("   - Function coverage")
	fmt.Println("   - Branch coverage")

	fmt.Println("\nNote: The actual tests and benchmarks are in calculator_test.go")
	fmt.Println("You can run them manually with:")
	fmt.Println("  go test ./calculator -v        # Run tests")
	fmt.Println("  go test ./calculator -bench=.  # Run benchmarks")
	fmt.Println("  go test ./calculator -cover    # Check coverage")

	// Run examples
	runTests()
	runBenchmarks()
	runCoverage()
}
