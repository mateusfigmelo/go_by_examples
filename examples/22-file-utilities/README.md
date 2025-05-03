# File Utilities in Go

This example demonstrates various file utility operations in Go, including line filters, file path manipulation, directory operations, temporary files, and the embed directive.

## Key Components

### 1. Line Filters
- Reading files line by line using `bufio.Scanner`
- Processing and transforming text content
- Efficient streaming of large files
- Pattern matching and text analysis

### 2. File Path Operations
- Path manipulation with `filepath` package
- Joining and splitting paths
- Getting absolute paths
- Cleaning and normalizing paths
- Working with file extensions and base names
- Cross-platform path handling

### 3. Directory Operations
- Creating directory structures
- Walking directory trees
- Filtering and processing files
- Recursive operations
- Directory listing and content analysis

### 4. Temporary Files and Directories
- Creating temporary files with `os.CreateTemp`
- Creating temporary directories with `os.MkdirTemp`
- Automatic cleanup with `defer`
- Pattern-based naming
- Safe concurrent usage

### 5. Embed Directive
- Embedding static files in binaries
- Using `//go:embed` directive
- Working with embedded file systems
- Template parsing from embedded files
- Configuration file management

## Best Practices

1. **File Operations**
   - Always close files using `defer`
   - Use buffered I/O for better performance
   - Handle errors appropriately
   - Clean up temporary resources

2. **Path Handling**
   - Use `filepath` package for cross-platform compatibility
   - Always clean user-provided paths
   - Use `filepath.Join` instead of string concatenation
   - Handle relative and absolute paths correctly

3. **Directory Operations**
   - Check permissions before operations
   - Handle errors during walks
   - Use appropriate buffer sizes
   - Implement proper cleanup

4. **Temporary Files**
   - Use unique names
   - Clean up after use
   - Handle concurrent access
   - Set appropriate permissions

5. **Embed Directive**
   - Organize embedded files logically
   - Use appropriate file patterns
   - Handle embedded resources efficiently
   - Consider binary size impact

## Memory Considerations
- Buffer sizes for file operations
- Memory usage during directory walks
- Embedded file size impact on binary
- Temporary file cleanup
- Resource management

## Common Patterns
1. Line Processing:
   - Text transformation
   - Pattern matching
   - Data extraction
   - Stream processing

2. Path Operations:
   - Configuration file locations
   - Plugin management
   - Resource loading
   - Cross-platform compatibility

3. Directory Management:
   - File organization
   - Backup operations
   - Content analysis
   - Resource cleanup

## Safety Considerations
1. Path traversal prevention
2. Permission handling
3. Resource cleanup
4. Error handling
5. Concurrent access

## Testing Considerations
1. Mock file systems
2. Temporary test directories
3. Platform-specific tests
4. Error condition testing
5. Cleanup verification

## Common Use Cases
1. Log file processing
2. Configuration management
3. Data import/export
4. File organization
5. Resource embedding

## Advanced Topics
1. Custom file systems
2. Virtual file systems
3. Network file systems
4. File watching
5. Extended attributes

## Running the Program
```bash
go run main.go
```

The program demonstrates:
1. Line-by-line file processing
2. File path manipulation
3. Directory tree walking
4. Temporary file handling
5. Embedded file usage

## Project Structure
```
22-file-utilities/
├── main.go           # Main program file
├── README.md         # Documentation
└── data/            # Embedded resources
    ├── config.txt    # Sample configuration
    └── template.html # Sample template
```

All temporary files and test directories are automatically cleaned up after the program runs. 
