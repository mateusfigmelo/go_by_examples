# 01 - Hello World

This is the first example in our Go programming series. It demonstrates the basic structure of a Go program.

## Concepts Covered

- Package declaration
- Importing standard libraries
- Main function
- Basic output using `fmt.Println`

## Code Explanation

```go
package main   // Declares the package name
import "fmt"   // Imports the fmt package for formatted I/O
func main() {  // Program entry point
    fmt.Println("Hello, World!")  // Prints text to console
}
```

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Expected Output
```
Hello, World!
```

## Key Takeaways

1. Every Go program must have a `package main` declaration if it's an executable program
2. The `main` function is the entry point of the program
3. The `fmt` package is commonly used for input/output operations
4. Go code is organized into packages for better modularity 
