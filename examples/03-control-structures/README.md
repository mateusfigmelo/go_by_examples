# 03 - Control Structures in Go

This example demonstrates all control structures available in Go, including loops, conditional statements, and switches. Each type of control structure is shown with practical examples and explanations.

## Concepts Covered

### Loop Types
1. Basic `for` loop (C-style)
2. While-style `for` loop
3. Infinite loop with `break`
4. Loop with `continue`
5. For-range loops (slices and maps)

### Conditional Statements
1. Basic `if-else`
2. `if` with initialization
3. `if-else if` chains

### Switch Statements
1. Basic switch
2. Switch with multiple cases
3. Switch with expression
4. Switch with fallthrough
5. Type switches (demonstrated in later examples)

### Special Features
1. Labels with break
2. Nested control structures
3. Loop control statements (`break`, `continue`)

## Key Points

1. Go only has the `for` loop, but it's flexible enough to handle all loop patterns
2. `switch` statements in Go don't fall through by default (unlike C/C++)
3. `if` statements can include an initialization statement
4. The `range` form of the `for` loop works with various data types
5. Labels can be used to break out of nested loops
6. Go doesn't have a `while` keyword, but `for` can be used to achieve the same effect

## Common Use Cases

### For Loops
- Basic iteration: `for i := 0; i < n; i++ { }`
- While-style: `for condition { }`
- Infinite loops: `for { }`
- Collection iteration: `for index, value := range collection { }`

### If Statements
- Simple conditions: `if x > y { }`
- With initialization: `if err := function(); err != nil { }`
- Multiple conditions: `if-else if-else` chains

### Switch Statements
- Value matching: `switch value { case x: ... }`
- Condition matching: `switch { case x > y: ... }`
- Type matching: `switch v.(type) { case string: ... }`

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Output Explanation

The program demonstrates each control structure with clear examples and output. Each section is numbered and labeled for easy reference. The output shows:

1. Different types of loops and their behavior
2. Conditional statement variations
3. Various switch statement use cases
4. Special control flow features like labels and nested structures

## Best Practices

1. Use the simplest loop form that fits your needs
2. Prefer `switch` over long `if-else` chains
3. Keep nesting levels manageable (prefer flat control structures)
4. Use meaningful labels when breaking out of nested loops
5. Consider readability when choosing between different control structures 
