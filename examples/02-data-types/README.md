# 02 - Data Types in Go

This example demonstrates all the basic data types available in Go and their usage. It covers numeric types, strings, booleans, and shows type inference, zero values, and type conversion.

## Concepts Covered

1. Basic Types:
   - Integers (signed and unsigned)
   - Floating-point numbers
   - Complex numbers
   - Booleans
   - Strings

2. Special Features:
   - Type inference
   - Zero values
   - Type conversion
   - Value ranges
   - Format printing

## Types Overview

### Integer Types
- Signed: `int8`, `int16`, `int32`, `int64`, `int`
- Unsigned: `uint8`, `uint16`, `uint32`, `uint64`, `uint`
- Special: `byte` (alias for `uint8`), `rune` (alias for `int32`)

### Floating-Point Types
- `float32` (single precision)
- `float64` (double precision)

### Complex Numbers
- `complex64` (32-bit real and imaginary components)
- `complex128` (64-bit real and imaginary components)

### Other Basic Types
- `bool` (true/false)
- `string` (UTF-8 encoded)

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Key Takeaways

1. Go is a statically typed language with type inference
2. Each type has a zero value (default value when not initialized)
3. Type conversion must be explicit in Go
4. The `int` and `uint` types are platform-dependent (32 or 64 bits)
5. Strings in Go are UTF-8 encoded by default
6. Complex numbers are built into the language
7. The `fmt` package provides rich formatting options for output

## Common Use Cases

- Use `int` for general integer values unless you need a specific size
- Use `float64` for floating-point calculations unless memory is critical
- Use `string` for text (UTF-8 encoded)
- Use `bool` for true/false conditions
- Use specific sized types (`int64`, etc.) when interfacing with external systems or when size matters 
