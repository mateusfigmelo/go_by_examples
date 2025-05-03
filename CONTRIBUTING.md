# Contributing to Go by Examples

Thank you for your interest in contributing to Go by Examples! This guide will help you understand how to structure your examples and ensure they meet our quality standards.

## Example Structure

Each example should follow this structure:

```
examples/XX-example-name/
├── README.md
├── main.go
└── main_test.go (required)
```

### README.md Requirements

Each example's README.md should include:

1. Title and brief description
2. Concepts covered
3. Code explanation
4. Running instructions
5. Expected output
6. Key takeaways
7. Test explanation (what aspects are being tested)

### Code Requirements

1. Each example should be self-contained in its directory
2. Main program should be in `main.go`
3. Tests MUST be included in `main_test.go`
4. Follow Go best practices and formatting guidelines

## Testing Requirements

Each example MUST include tests that:

1. Verify the core functionality being demonstrated
2. Include both positive and negative test cases where applicable
3. Mock external dependencies (files, network, etc.) when needed
4. Use table-driven tests for multiple test cases
5. Have clear test names and failure messages
6. Utilize the provided test utilities (`testutil` package) where applicable

### Testing Examples

Here's how to structure your tests based on different example types:

1. **For Function Examples:**
```go
// main.go
package main

func add(a, b int) int {
    return a + b
}

func main() {
    // Example usage
    add(1, 2)
}

// main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"negative numbers", -1, -2, -3},
        {"zero values", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := add(tt.a, tt.b); got != tt.expected {
                t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

2. **For Console Output Examples:**
```go
// Using our testutil package to capture and test output
import (
    "testing"
    "github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestOutputFunction(t *testing.T) {
    output := testutil.CaptureOutput(func() {
        outputFunction()
    })
    
    expected := "Expected output"
    if !testutil.ContainsOutput(output, expected) {
        t.Errorf("Expected output to contain %q, got %q", expected, output)
    }
}
```

3. **For File/Network Examples:**
```go
// Use test utilities and mocks
func TestFileProcessing(t *testing.T) {
    // Create a temporary test file
    tmpFile, err := os.CreateTemp("", "test-*.txt")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpFile.Name())

    // Test file operations
    // ...
}
```

## Quality Guidelines

1. Code must be formatted using `gofmt`
2. Examples should be clear and focused on a single concept
3. Include meaningful comments
4. Add proper error handling (check all error return values)
5. ALL examples must have comprehensive tests

## Automated Checks

Our GitHub Actions workflow performs the following checks:

1. **Formatting**: Verifies code is properly formatted with `gofmt`
2. **Linting**: Runs `golangci-lint` to check for common issues
3. **Static analysis**: Uses `staticcheck` to find potential bugs
4. **Test verification**: Ensures all examples have test files
5. **Test execution**: Runs all tests and generates coverage reports

## Pull Request Process

1. Fork the repository and create a new branch for your example
2. Follow the example structure above
3. Ensure your code passes all automated checks:
   - Code is properly formatted (`gofmt -w .`)
   - All linter issues are resolved
   - All tests pass
4. Submit a PR with a clear description of your example and its tests

## Running Tests Locally

Before submitting a PR, please run:

```bash
# Format code
gofmt -w .

# Run linters (install if needed)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run ./...

go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...

# Run tests for all examples
go run main.go

# Run tests for your specific example
go run main.go XX-your-example

# Run tests with coverage
go run main.go -c XX-your-example

# Run tests with verbose output
go run main.go -v XX-your-example
```

## Questions?

If you have any questions, feel free to open an issue for discussion. 
