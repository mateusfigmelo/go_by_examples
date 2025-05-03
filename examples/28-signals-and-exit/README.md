# Signal and Exit Handling in Go

This example demonstrates how to handle system signals and program exit scenarios in Go, including graceful shutdown, panic recovery, and worker management.

## Project Structure
```
28-signals-and-exit/
├── main.go    # Main program with signal and exit handling examples
└── README.md  # Documentation
```

## Key Components

### 1. Signal Handling
- SIGINT (Ctrl+C) handling
- SIGTERM handling
- SIGHUP for configuration reload
- SIGUSR1 for status reports
- Graceful shutdown

### 2. Exit Management
- Cleanup handlers
- Panic recovery
- Exit codes
- Resource cleanup
- Graceful worker shutdown

### 3. Worker Management
- Worker pool implementation
- Context-based cancellation
- Thread-safe operations
- Worker lifecycle management

## Running the Example

1. **Basic Run**
   ```bash
   go run main.go
   ```

2. **Test Different Scenarios**
   ```bash
   # Test panic recovery
   go run main.go panic

   # Test error exit
   go run main.go error

   # Send signals (Unix/Linux)
   kill -HUP <pid>   # Reload workers
   kill -USR1 <pid>  # Status report
   kill -TERM <pid>  # Graceful shutdown
   ```

## Features Demonstrated

1. **Signal Handling**
   - Multiple signal types
   - Custom signal actions
   - Graceful shutdown
   - Worker management

2. **Exit Scenarios**
   - Normal exit (code 0)
   - Error exit (code 1)
   - Panic recovery
   - Resource cleanup

3. **Worker Management**
   - Context-based cancellation
   - Thread-safe operations
   - Worker lifecycle
   - Pool management

## Best Practices

1. **Signal Handling**
   - Use buffered channels
   - Handle multiple signals
   - Implement graceful shutdown
   - Clean up resources

2. **Exit Management**
   - Use defer for cleanup
   - Recover from panics
   - Use appropriate exit codes
   - Log exit reasons

3. **Resource Management**
   - Use mutexes for thread safety
   - Implement proper cleanup
   - Handle goroutine shutdown
   - Prevent resource leaks

## Common Patterns

1. **Graceful Shutdown**
   ```go
   sigChan := make(chan os.Signal, 1)
   signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
   <-sigChan
   // Perform cleanup
   ```

2. **Panic Recovery**
   ```go
   defer func() {
       if r := recover(); r != nil {
           log.Printf("Recovered: %v", r)
           os.Exit(1)
       }
   }()
   ```

3. **Worker Management**
   ```go
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   // Start workers with context
   ```

## Additional Resources
- [Signal Package Documentation](https://golang.org/pkg/os/signal/)
- [OS Package Documentation](https://golang.org/pkg/os/)
- [Context Package Documentation](https://golang.org/pkg/context/)
- [Sync Package Documentation](https://golang.org/pkg/sync/) 
