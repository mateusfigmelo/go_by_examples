# Go By Examples

<div align="center">
  <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" width="200" alt="GoLang">
  <br><br>

  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](https://github.com/mateusfigmelo/go_by_examples/pulls)
  [![Stars](https://img.shields.io/github/stars/mateusfigmelo/go_by_examples?style=social)](https://github.com/mateusfigmelo/go_by_examples/stargazers)
  
  [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example)](https://pkg.go.dev/golang.org/x/example)
  [![Go Report Card](https://goreportcard.com/badge/github.com/mateusfigmelo/go_by_examples)](https://goreportcard.com/report/github.com/mateusfigmelo/go_by_examples)
  [![Last Commit](https://img.shields.io/github/last-commit/mateusfigmelo/go_by_examples)](https://github.com/mateusfigmelo/go_by_examples/commits/main)
  [![Examples](https://img.shields.io/badge/examples-28-00ADD8)](https://github.com/mateusfigmelo/go_by_examples)

  <h3>A comprehensive collection of Go programming examples</h3>
  <p>Learn Go by example: from basic concepts to advanced patterns and best practices</p>
</div>

## Examples

1. [Hello World](./examples/01-hello-world) - Basic program structure
2. [Data Types](./examples/02-data-types) - Types, variables, and constants
3. [Control Structures](./examples/03-control-structures) - Flow control and branching
4. [Collections](./examples/04-collections) - Arrays, slices, and maps
5. [Functions](./examples/05-functions) - Function declarations and usage
6. [Types and Pointers](./examples/06-types-and-pointers) - Memory management
7. [Structs and Interfaces](./examples/07-structs-and-interfaces) - Custom types and abstraction
8. [Generics and Embedding](./examples/08-generics-and-embedding) - Generic programming and composition
9. [Error Handling](./examples/09-error-handling) - Error types and handling
10. [Concurrency](./examples/10-concurrency) - Goroutines and basic concurrency
11. [Advanced Channels](./examples/11-advanced-channels) - Channel patterns and usage
12. [Worker Pools](./examples/12-worker-pools) - Concurrent worker patterns
13. [Sync Primitives](./examples/13-sync-primitives) - Mutexes and synchronization
14. [Sorting](./examples/14-sorting) - Sorting algorithms and interfaces
15. [Panic, Defer, Recover](./examples/15-panic-defer-recover) - Error flow control
16. [Strings and Templates](./examples/16-strings-and-templates) - String manipulation
17. [JSON and XML](./examples/17-json-xml) - Data serialization
18. [Time Handling](./examples/18-time-handling) - Time operations
19. [Parsing and Random](./examples/19-parsing-and-random) - Text parsing and randomization
20. [Hashes and Encoding](./examples/20-hashes-and-encoding) - Cryptography basics
21. [File Operations](./examples/21-file-operations) - File system handling
22. [File Utilities](./examples/22-file-utilities) - Advanced file operations
23. [Testing and Benchmarking](./examples/23-testing-benchmarking) - Advanced testing
24. [CLI](./examples/24-cli) - Command-line interfaces
25. [Environment and Logging](./examples/25-env-and-logging) - Configuration and logs
26. [HTTP Client and Server](./examples/26-http) - HTTP client/server
27. [Context and Processes](./examples/27-context-and-processes) - Process management
28. [Signals and Exit](./examples/28-signals-and-exit) - Signal handling

## Prerequisites

- Go 1.21 or higher (some examples use newer features)
- Basic understanding of Go programming concepts
- Text editor or IDE
- Terminal/Command Prompt

## Running Examples

Each example directory contains:
- Source code files (`.go`)
- README with explanations and instructions
- Tests where applicable
- Any necessary resource files

To run an example:
```bash
cd examples/<example-directory>
go run main.go
```

## Testing

Each example includes tests that demonstrate proper usage and validate functionality. We use Go's built-in testing framework with some extended features.

### Running Tests

You can run tests using our Go test runner:

```bash
# Run all tests
go run main.go

# Run tests for a specific example
go run main.go 12-worker-pools

# Run tests with coverage
go run main.go -c

# Get verbose test output
go run main.go -v

# See all options
go run main.go -h
```

### Test Coverage

Coverage reports are generated in HTML format when running tests with the `-c` or `--coverage` flag. These reports are saved in the `coverage/` directory.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-example`)
3. Make your changes and add appropriate tests
4. Ensure all tests pass (`go run main.go`)
5. Commit your changes (`git commit -m 'Add a new example'`)
6. Push to the branch (`git push origin feature/amazing-example`)
7. Create a Pull Request

## Best Practices

Each example follows these principles:
- Clear and documented code
- Proper error handling
- Comprehensive unit tests 
- Table-driven testing where appropriate
- Consistent formatting (using `go fmt`)
- Idiomatic Go patterns
- Performance considerations
- Security best practices

### Testing Best Practices

- Write tests for all exported functions
- Use table-driven tests for multiple scenarios
- Aim for high code coverage (at least 80%)
- Test both happy paths and error cases
- Keep tests independent and idempotent
- Use mocks or stubs for external dependencies
- Include benchmarks for performance-critical code
- Write clear failure messages

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
