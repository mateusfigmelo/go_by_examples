# 07 - Structs, Methods, and Interfaces in Go

This example demonstrates Go's object-oriented features through structs, methods, interfaces, and type-based enums. While Go isn't traditionally object-oriented, it provides powerful features for organizing and structuring code.

## Concepts Covered

### 1. Structs
- Basic struct definition
- Nested structs
- Struct tags
- Struct embedding
- Field visibility (exported vs unexported)
- Struct initialization
- Anonymous structs

### 2. Methods
- Method receivers
- Pointer receivers vs value receivers
- Method sets
- Method chaining
- Method overriding
- String() method implementation

### 3. Interfaces
- Interface definition
- Interface implementation (implicit)
- Interface composition
- Empty interface
- Type assertions
- Type switches
- Interface satisfaction

### 4. Type-Based Enums
- Iota usage
- Enum methods
- String representation
- Type safety
- Custom behavior
- Enum patterns

## Key Components

### Geometric Shapes Example
1. Shape Interface
   - Area() method
   - Perimeter() method
   - String() method

2. Concrete Types
   - Point struct
   - Circle struct
   - Rectangle struct

3. Method Implementations
   - Area calculations
   - Perimeter calculations
   - String representations

### Person Example
1. Nested Structs
   - Address within Person
   - Struct tags
   - JSON annotations

2. Methods
   - Birthday() method
   - String() method

### Interface Composition
1. Logger Interface
2. Writer Interface
3. LogWriter Combined Interface
4. ConsoleLogger Implementation

### Enum Types
1. Direction Enum
   - North, East, South, West
   - String() method

2. Status Enum
   - Pending, Active, Inactive
   - Custom behavior (IsActive)

## Best Practices

### Struct Design
1. Keep structs focused and cohesive
2. Use embedding for composition
3. Consider visibility requirements
4. Document fields and methods
5. Use meaningful field names

### Method Design
1. Choose receiver type carefully
2. Document method behavior
3. Maintain consistency
4. Consider performance implications
5. Follow naming conventions

### Interface Design
1. Keep interfaces small
2. Design for behavior
3. Use composition
4. Follow interface segregation
5. Document expectations

### Enum Design
1. Use iota appropriately
2. Provide String() method
3. Consider type safety
4. Add domain-specific behavior
5. Document valid values

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Memory and Performance Considerations

### Struct Considerations
1. Size and alignment
2. Value vs pointer semantics
3. Copy overhead
4. Memory layout
5. Field ordering

### Method Considerations
1. Receiver type impact
2. Escape analysis
3. Method call overhead
4. Interface method calls
5. Inlining possibilities

### Interface Considerations
1. Dynamic dispatch overhead
2. Type assertion costs
3. Interface pollution
4. Memory layout
5. Runtime type information

## Common Patterns and Idioms

### Struct Patterns
1. Builder pattern
2. Factory methods
3. Option pattern
4. Immutable structs
5. Validation methods

### Interface Patterns
1. io.Reader/io.Writer
2. Stringer interface
3. Error interface
4. Interface composition
5. Interface wrapping

### Method Patterns
1. Functional options
2. Method chaining
3. Error handling
4. Fluent interfaces
5. Constructor methods

## Error Handling and Safety

1. Nil pointer checks
2. Type assertion safety
3. Interface compliance
4. Error propagation
5. Panic recovery

## Testing Considerations

1. Table-driven tests
2. Interface mocking
3. Struct comparison
4. Method testing
5. Interface testing 
