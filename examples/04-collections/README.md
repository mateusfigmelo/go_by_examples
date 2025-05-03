# 04 - Collections in Go

This example demonstrates Go's collection types: arrays, slices, and maps. It covers their creation, manipulation, and common operations with practical examples.

## Concepts Covered

### 1. Arrays
- Fixed-size sequences of elements
- Declaration and initialization
- Zero-valued arrays
- Array literals
- Multi-dimensional arrays
- Array operations

### 2. Slices
- Dynamic-size sequences
- Creation methods:
  - Slice literals
  - `make` function
  - Slicing from arrays
- Operations:
  - Append
  - Copy
  - Slicing
  - Length and capacity

### 3. Maps
- Key-value pairs
- Creation methods:
  - `make` function
  - Map literals
- Operations:
  - Adding/updating entries
  - Accessing values
  - Checking existence
  - Deletion
  - Iteration

### 4. Advanced Operations
- Sorting slices
- Working with keys and values
- Type-specific operations
- Common patterns and idioms

## Key Points

### Arrays
1. Fixed length, part of type
2. Zero-based indexing
3. Passed by value to functions
4. Can be multi-dimensional
5. Size must be known at compile time

### Slices
1. Dynamic length, built on arrays
2. Reference type
3. Can grow with `append`
4. Has length and capacity
5. Common operations: append, copy, slice

### Maps
1. Unordered key-value pairs
2. Reference type
3. Keys must be comparable
4. Thread-unsafe by default
5. Common operations: insert, delete, lookup

## Common Use Cases

### Arrays
- Fixed-size buffers
- Small sequences with known size
- Matrix operations
- Performance-critical code

### Slices
- Dynamic lists
- Buffer management
- Stack/queue implementations
- Processing sequences

### Maps
- Lookup tables
- Caching
- Counting/frequency tracking
- De-duplication
- Graph representations

## Running the Program

From this directory, run:
```bash
go run main.go
```

## Best Practices

1. Arrays
   - Use when size is fixed and known
   - Consider slices for most other cases
   - Be aware of copying behavior

2. Slices
   - Preallocate with make() when size is known
   - Use append() for dynamic growth
   - Watch for capacity changes
   - Be careful with re-slicing large arrays

3. Maps
   - Always check for existence when needed
   - Use two-value assignment for safe access
   - Delete keys you no longer need
   - Consider sync.Map for concurrent access

## Memory Considerations

1. Arrays
   - Stack allocated when small
   - Pass by value (copies)
   - Size affects performance

2. Slices
   - Reference type (header is small)
   - Underlying array can be shared
   - Growth may cause reallocation

3. Maps
   - Reference type
   - Can grow dynamically
   - Memory overhead per entry 
