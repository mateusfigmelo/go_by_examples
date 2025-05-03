# 05 - Functions in Go

This example demonstrates the various ways to define and use functions in Go, including advanced features like multiple return values, variadic functions, closures, and recursion.

## Concepts Covered

### 1. Basic Functions
- Function declaration and definition
- Parameters and return values
- Function signatures
- Basic function calls

### 2. Advanced Function Features
- Multiple return values
- Named return values
- Variadic functions
- Function types
- Anonymous functions
- Functions as values

### 3. Closures
- Function literals
- Capturing variables
- Factory functions
- Stateful functions

### 4. Recursion
- Basic recursion
- Recursive functions with error handling
- Common recursive algorithms
- Stack considerations

### 5. Special Features
- Deferred function calls
- Function parameters
- Function return values
- Value vs pointer receivers

## Key Points

### Function Declaration
1. Functions are first-class citizens
2. Can return multiple values
3. Parameters are pass-by-value
4. Can have named return values
5. Support variadic parameters

### Closures
1. Can capture variables from outer scope
2. Maintain reference to captured variables
3. Useful for stateful functions
4. Common in callbacks and goroutines

### Recursion
1. Functions can call themselves
2. Need base case to terminate
3. Stack space consideration
4. Useful for tree/graph algorithms

### Special Features
1. Defer statements for cleanup
2. Error handling patterns
3. Function type declarations
4. Method declarations

## Common Use Cases

### Multiple Return Values
- Error handling
- Compound calculations
- Optional values
- State and status returns

### Variadic Functions
- String formatting
- Collection operations
- Optional parameters
- Logging functions

### Closures
- Event handlers
- Middleware
- Partial application
- State encapsulation

### Recursion
- Tree traversal
- Mathematical calculations
- Problem subdivision
- Graph algorithms

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Best Practices

1. Function Design
   - Keep functions focused and small
   - Use meaningful parameter names
   - Document complex functions
   - Return errors as second value

2. Error Handling
   - Check for errors immediately
   - Propagate errors up
   - Provide context in errors
   - Use custom error types when needed

3. Closures
   - Be careful with variable capture
   - Document closure behavior
   - Consider memory implications
   - Use for clear encapsulation

4. Recursion
   - Ensure termination conditions
   - Consider stack limits
   - Use iteration for simple loops
   - Document recursive behavior

## Memory Considerations

1. Function Calls
   - Parameters are copied
   - Large structs should use pointers
   - Stack vs heap allocation
   - Garbage collection impact

2. Closures
   - Captured variables stay alive
   - Memory leaks possible
   - Reference cycles
   - Cleanup considerations

3. Recursion
   - Stack space usage
   - Tail call optimization
   - Memory usage in deep recursion
   - Alternative iterative solutions 
