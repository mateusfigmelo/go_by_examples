# Cryptographic Hashes and Encoding in Go

This example demonstrates cryptographic hash functions and various encoding methods in Go, including secure hashing algorithms, HMAC, and different encoding formats.

## Key Components

### 1. Hash Functions
- SHA-256 implementation
- SHA-512 implementation
- SHA-1 (with security warnings)
- MD5 (with security warnings)
- Streaming hash computation

### 2. HMAC (Hash-based Message Authentication Code)
- HMAC-SHA256 implementation
- HMAC-SHA512 implementation
- Message verification
- Security considerations

### 3. Common Hashing Patterns
- Multiple item hashing
- Password hashing with salt
- File hashing
- Stream processing

### 4. Base64 Encoding
- Standard Base64
- URL-safe Base64
- Raw Base64 (no padding)
- Custom Base64 alphabets
- Encoding/decoding operations

### 5. Base32 Encoding
- Standard Base32
- Hex Base32
- Custom Base32 alphabets
- Padding options

### 6. Hex Encoding
- Standard hex encoding
- Hex dumps
- Binary-to-hex conversion

## Best Practices

1. **Cryptographic Hashing**
   - Use SHA-256 or SHA-512 for new applications
   - Avoid MD5 and SHA-1 for security-critical operations
   - Always handle hash computation errors
   - Use streaming for large data

2. **HMAC Usage**
   - Keep keys secure and private
   - Use cryptographically secure key generation
   - Implement constant-time comparisons
   - Handle all error cases

3. **Encoding**
   - Choose appropriate encoding for the use case
   - Handle padding correctly
   - Validate input before decoding
   - Consider URL safety when needed

## Memory Considerations

1. **Hash Computation**
   - Buffer management for streaming
   - Memory usage in hash state
   - Output size considerations
   - Cleanup of sensitive data

2. **HMAC Operations**
   - Key storage security
   - Memory cleanup after use
   - Buffer reuse strategies
   - Resource management

3. **Encoding Operations**
   - Output size calculations
   - Buffer pre-allocation
   - Memory efficiency in conversions
   - Streaming for large data

## Common Patterns

1. **File Hashing**
   ```go
   hash := sha256.New()
   if _, err := io.Copy(hash, file); err != nil {
       // Handle error
   }
   sum := hash.Sum(nil)
   ```

2. **Password Hashing**
   ```go
   hash := sha256.New()
   hash.Write(salt)
   hash.Write(password)
   hashedPassword := hash.Sum(nil)
   ```

3. **Base64 Encoding**
   ```go
   encoded := base64.StdEncoding.EncodeToString(data)
   decoded, err := base64.StdEncoding.DecodeString(encoded)
   ```

## Safety Considerations

1. **Cryptographic Security**
   - Use cryptographically secure algorithms
   - Keep keys and sensitive data secure
   - Implement proper error handling
   - Consider timing attacks

2. **Input Validation**
   - Validate input sizes
   - Check for malformed data
   - Handle decoding errors
   - Sanitize output when needed

3. **Resource Management**
   - Clean up sensitive data
   - Handle large inputs safely
   - Manage memory efficiently
   - Close resources properly

## Testing Considerations

1. **Hash Testing**
   - Test known input/output pairs
   - Test streaming operations
   - Test error conditions
   - Test with different sizes

2. **HMAC Testing**
   - Test with various keys
   - Test verification
   - Test tampering detection
   - Test error cases

3. **Encoding Testing**
   - Test encoding/decoding pairs
   - Test special characters
   - Test padding cases
   - Test error handling

## Common Use Cases

1. **Data Integrity**
   - File checksums
   - Message verification
   - Content addressing
   - Cache invalidation

2. **Security**
   - Password hashing
   - Digital signatures
   - Token generation
   - Message authentication

3. **Data Encoding**
   - API data transfer
   - URL-safe data
   - Binary data representation
   - Configuration storage

## Advanced Topics

1. **Custom Hash Functions**
   - Implementing hash.Hash interface
   - Custom HMAC implementations
   - Performance optimization
   - Security considerations

2. **Advanced Encoding**
   - Custom encodings
   - Streaming encoders
   - Performance tuning
   - Error recovery

3. **Security Considerations**
   - Side-channel attacks
   - Timing attacks
   - Memory safety
   - Key management

## Running the Program

```bash
go run main.go
```

The program demonstrates:
- Various hash functions
- HMAC operations
- Common hashing patterns
- Different encoding formats
- Error handling and validation

## Further Reading

- [crypto Package Documentation](https://golang.org/pkg/crypto/)
- [encoding Package Documentation](https://golang.org/pkg/encoding/)
- [HMAC: Keyed-Hashing for Message Authentication](https://tools.ietf.org/html/rfc2104)
- [Base64 Encoding](https://tools.ietf.org/html/rfc4648) 
