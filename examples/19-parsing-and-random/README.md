# Random Numbers, Number Parsing, and URL Parsing in Go

This example demonstrates random number generation, number parsing, and URL parsing in Go, including secure random number generation, string-to-number conversions, and URL manipulation.

## Key Components

### 1. Random Number Generation
- Cryptographically secure random numbers using `crypto/rand`
- Random byte generation
- Random string generation
- Best practices for secure randomization

### 2. Number Parsing
- String to integer conversion
- Base-specific number parsing (hex, binary)
- Float parsing
- Boolean parsing
- Error handling for invalid inputs
- Number formatting and base conversion

### 3. URL Parsing
- Basic URL parsing
- URL component access
- Query parameter handling
- URL building
- Path escaping/unescaping
- Relative URL resolution

### 4. Advanced URL Operations
- URL normalization
- URL validation
- Authentication information handling
- Complex query parameter manipulation

## Best Practices

1. **Random Number Generation**
   - Use `crypto/rand` for secure random numbers
   - Never use `math/rand` for security-critical operations
   - Always check errors from random number generation
   - Use appropriate buffer sizes for random bytes

2. **Number Parsing**
   - Always handle parsing errors
   - Use appropriate bit sizes for integers
   - Validate input ranges
   - Consider locale-specific formatting

3. **URL Handling**
   - Validate URLs before processing
   - Always escape user input
   - Handle query parameters safely
   - Consider URL normalization for comparisons

## Memory Considerations

1. **Random Number Generation**
   - Buffer allocation for random bytes
   - String builder efficiency for random strings
   - Memory usage in cryptographic operations

2. **Number Parsing**
   - String conversion overhead
   - Buffer management for large numbers
   - Memory efficiency in base conversions

3. **URL Parsing**
   - URL structure memory layout
   - Query parameter storage efficiency
   - String manipulation overhead

## Common Patterns

1. **Random Number Generation**
   ```go
   n, err := rand.Int(rand.Reader, big.NewInt(100))
   if err != nil {
       // Handle error
   }
   ```

2. **Number Parsing**
   ```go
   num, err := strconv.Atoi(str)
   if err != nil {
       // Handle invalid number
   }
   ```

3. **URL Parsing**
   ```go
   u, err := url.Parse(rawURL)
   if err != nil {
       // Handle invalid URL
   }
   ```

## Safety Considerations

1. **Random Number Security**
   - Use cryptographic random for security
   - Verify random number ranges
   - Handle entropy exhaustion
   - Consider timing attacks

2. **Number Parsing Safety**
   - Validate input ranges
   - Handle overflow conditions
   - Consider locale-specific issues
   - Validate decimal places

3. **URL Safety**
   - Validate URL schemes
   - Escape user input
   - Handle malformed URLs
   - Consider security implications of redirects

## Testing Considerations

1. **Random Number Testing**
   - Test error conditions
   - Verify distribution (where applicable)
   - Mock random source for reproducible tests
   - Test buffer handling

2. **Number Parsing Testing**
   - Test edge cases
   - Test invalid inputs
   - Test different bases
   - Test locale-specific formats

3. **URL Testing**
   - Test malformed URLs
   - Test special characters
   - Test query parameter handling
   - Test relative URL resolution

## Common Use Cases

1. **Random Numbers**
   - Generating session IDs
   - Creating random tokens
   - Cryptographic operations
   - Testing and simulation

2. **Number Parsing**
   - User input processing
   - Configuration parsing
   - Data conversion
   - Scientific calculations

3. **URL Parsing**
   - Web application routing
   - API client implementation
   - Link validation
   - Web scraping

## Advanced Topics

1. **Custom Random Sources**
   - Implementing custom random sources
   - Seeding strategies
   - Distribution algorithms
   - Performance optimization

2. **Advanced Number Formatting**
   - Custom number formats
   - Locale-specific formatting
   - Scientific notation
   - Custom base conversion

3. **URL Processing**
   - Custom URL schemes
   - URL normalization rules
   - Internationalized URLs
   - URL pattern matching

## Running the Program

```bash
go run main.go
```

The program demonstrates:
- Secure random number generation
- Random string creation
- Number parsing and formatting
- Basic and advanced URL operations
- Error handling and validation

## Further Reading

- [crypto/rand Package Documentation](https://golang.org/pkg/crypto/rand/)
- [strconv Package Documentation](https://golang.org/pkg/strconv/)
- [net/url Package Documentation](https://golang.org/pkg/net/url/)
- [RFC 3986 - URI Generic Syntax](https://tools.ietf.org/html/rfc3986) 
