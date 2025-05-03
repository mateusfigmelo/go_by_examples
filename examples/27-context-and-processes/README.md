# Context and Process Management in Go

This example demonstrates how to work with contexts and manage processes in Go, including spawning processes, executing commands, and handling process lifecycle.

## Project Structure
```
27-context-and-processes/
├── main.go    # Main program with context and process examples
└── README.md  # Documentation
```

## Key Components

### 1. Context Management
- Context creation and cancellation
- Timeout and deadline handling
- Context propagation
- Error handling with context
- Context values and metadata

### 2. Process Management
- Process spawning and execution
- Process monitoring
- Output capturing
- Error handling
- Process termination
- Cross-platform compatibility

### 3. Process Manager
- Concurrent process handling
- Process lifecycle management
- Resource cleanup
- Signal handling
- Error propagation

## Features Demonstrated

### 1. Context Usage
```go
// Creating a context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// Creating a context with cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Using context with commands
cmd := exec.CommandContext(ctx, command, args...)
```

### 2. Process Spawning
```go
// Starting a process with context
cmd := exec.CommandContext(ctx, command, args...)
stdout, _ := cmd.StdoutPipe()
stderr, _ := cmd.StderrPipe()

if err := cmd.Start(); err != nil {
    return fmt.Errorf("failed to start process: %v", err)
}

// Handling process completion
err := cmd.Wait()
if err != nil {
    if ctx.Err() == context.Canceled {
        log.Printf("Process canceled by context")
    } else if _, ok := err.(*exec.ExitError); ok {
        log.Printf("Process exited with error")
    }
}
```

### 3. Process Output Handling
```go
// Capturing output with proper error handling
stdout, err := cmd.StdoutPipe()
if err != nil {
    return fmt.Errorf("failed to create stdout pipe: %v", err)
}

// Reading output with context
go func() {
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        log.Printf("[%s] %s", name, scanner.Text())
    }
}()
```

## Best Practices

### 1. Context Management
- Always call cancel() when done
- Use appropriate timeout values
- Propagate context through function calls
- Handle context cancellation gracefully
- Clean up resources on cancellation
- Provide clear timeout messages

### 2. Process Management
- Use process groups for related processes
- Handle process termination signals
- Clean up zombie processes
- Set appropriate timeouts
- Handle output buffering
- Use appropriate error handling
- Implement graceful shutdown
- Use descriptive process names

### 3. Resource Management
- Close pipes and files
- Release system resources
- Handle process cleanup
- Manage memory usage
- Use appropriate buffer sizes
- Implement mutex protection for shared resources
- Clean up processes on program exit

## Common Patterns

### 1. Process Manager
```go
type ProcessManager struct {
    processes map[string]*exec.Cmd
    mu        sync.RWMutex
}

// Managing processes with context
pm.StartProcess(ctx, "process1", command, args...)
pm.StopProcess("process1")
pm.StopAll()
```

### 2. Output Multiplexing
```go
// Handling multiple output streams with context
go func() {
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        log.Printf("[%s] %s", name, scanner.Text())
    }
}()

go func() {
    scanner := bufio.NewScanner(stderr)
    for scanner.Scan() {
        log.Printf("[%s][ERROR] %s", name, scanner.Text())
    }
}()
```

### 3. Graceful Shutdown
```go
// Setting up signal handling
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

// Handling shutdown
select {
case <-sigChan:
    log.Println("Received interrupt signal")
    cancel()
    pm.StopAll()
case <-ctx.Done():
    log.Println("Context canceled")
default:
    pm.StopAll()
}
```

## Running the Example

1. **Start the Program**
   ```bash
   go run main.go
   ```

2. **Example Output**
   ```
   === Context and Process Management Example ===

   === Demonstrating Context Timeout ===
   Starting context timeout demonstration...
   Starting long-running command (sleep 5)...
   Command was terminated due to context timeout (after 2 seconds)
   Context timeout demonstration completed

   === Demonstrating Process Spawning ===
   Starting process spawning demonstration...
   Starting process: counter1
   Starting process: counter2
   [counter1] Counter 1: 1
   [counter2] Counter 2: 1
   ...
   Stopping counter1 process...
   [counter1] Process exited with error
   [counter2] Process completed successfully

   === Demonstrating Process Execution ===
   Starting process execution demonstration...
   Running 'go version'...
   Go version output: go1.21.0
   ```

3. **Interrupt Handling**
   - Press Ctrl+C to trigger graceful shutdown
   - All processes will be terminated with proper cleanup
   - Resources will be released systematically
   - Proper error messages will be logged

## Common Issues and Solutions

### 1. Process Leaks
- Always wait for processes to complete
- Use context for automatic cancellation
- Clean up zombie processes
- Handle all error cases
- Implement proper process tracking
- Use mutex protection for process management

### 2. Resource Exhaustion
- Limit number of concurrent processes
- Set appropriate timeouts
- Monitor resource usage
- Clean up unused resources
- Implement proper error handling
- Use buffered channels for signal handling

### 3. Cross-Platform Compatibility
- Use platform-specific commands carefully
- Handle path separators
- Consider process signals
- Test on all target platforms
- Use runtime.GOOS for OS-specific logic
- Handle Windows/Unix command differences

## Additional Resources
- [context Package Documentation](https://golang.org/pkg/context/)
- [os/exec Package Documentation](https://golang.org/pkg/os/exec/)
- [Process Management in Go](https://golang.org/pkg/os/#Process)
- [Signal Handling in Go](https://golang.org/pkg/os/signal/) 
