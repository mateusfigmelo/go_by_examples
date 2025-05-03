# Testing and Benchmarking in Go

This example demonstrates comprehensive testing and benchmarking in Go, including unit tests, table-driven tests, benchmarks, and code coverage analysis.

## Key Components

### 1. Unit Testing
- Basic test functions
- Table-driven tests
- Subtests for better organization
- Error testing
- Test fixtures and setup/teardown
- Testing edge cases
- Example functions for documentation

### 2. Benchmarking
- Performance benchmarks
- Memory allocation measurements
- Comparative benchmarks
- Benchmark fixtures
- Custom benchmark timing
- Resource usage analysis

### 3. Coverage Analysis
- Code coverage reporting
- Function coverage
- Branch coverage
- Coverage profiles
- HTML coverage reports

## Project Structure
```
23-testing-benchmarking/
├── main.go                 # Example runner
├── README.md              # Documentation
└── calculator/            # Package being tested
    ├── calculator.go      # Main implementation
    └── calculator_test.go # Tests and benchmarks
```

## Test Types

### 1. Basic Tests
```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

### 2. Table-Driven Tests
```go
func TestMultiply(t *testing.T) {
    tests := []struct {
        x, y, want int
    }{
        {2, 3, 6},
        {-2, 3, -6},
        {0, 5, 0},
    }
    for _, tt := range tests {
        if got := Multiply(tt.x, tt.y); got != tt.want {
            t.Errorf("Multiply(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.want)
        }
    }
}
```

### 3. Benchmarks
```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(20)
    }
}
```

## Best Practices

1. **Test Organization**
   - Use table-driven tests for multiple test cases
   - Group related tests using subtests
   - Keep test files alongside implementation
   - Use meaningful test names

2. **Test Coverage**
   - Aim for high test coverage
   - Test edge cases and error conditions
   - Test both valid and invalid inputs
   - Test concurrent operations

3. **Benchmarking**
   - Reset timers when necessary
   - Use realistic data sizes
   - Compare alternative implementations
   - Consider memory allocations

4. **Documentation**
   - Include example functions
   - Document test helpers
   - Explain complex test cases
   - Document benchmark methodology

## Running Tests

### Basic Testing
```bash
go test ./...                 # Test all packages
go test -v ./...             # Verbose output
go test -run TestAdd         # Run specific test
```

### Benchmarking
```bash
go test -bench=.             # Run all benchmarks
go test -bench=Add          # Run specific benchmark
go test -bench=. -benchmem  # Include memory stats
```

### Coverage
```bash
go test -cover              # Show coverage
go test -coverprofile=c.out # Generate coverage profile
go tool cover -html=c.out   # View HTML coverage report
```

## Common Patterns

1. **Test Helpers**
   - Setup/teardown functions
   - Test data generators
   - Custom assertions
   - Error checkers

2. **Mocking**
   - Interface-based design
   - Mock implementations
   - Stub functions
   - Dependency injection

3. **Parallel Testing**
   - t.Parallel()
   - Race condition detection
   - Concurrent test execution
   - Resource cleanup

## Advanced Topics

1. **Testing Tools**
   - Race detector
   - Fuzzing
   - Profiling
   - Custom test flags

2. **Performance Testing**
   - CPU profiling
   - Memory profiling
   - Goroutine profiling
   - Trace analysis

3. **Integration Testing**
   - External dependencies
   - Database testing
   - Network testing
   - API testing

## Running the Example
```bash
go run main.go
```

This will:
1. Run all unit tests
2. Execute benchmarks
3. Generate and display coverage report

## Additional Resources
- [Go Testing Package Documentation](https://golang.org/pkg/testing/)
- [Go Blog: Using Subtests and Sub-benchmarks](https://blog.golang.org/subtests)
- [Go Blog: Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
- [Go Blog: Coverage Profiling](https://blog.golang.org/cover) 
