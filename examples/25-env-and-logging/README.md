# Environment Variables and Logging in Go

This example demonstrates best practices for handling environment variables and implementing logging in Go applications.

## Project Structure
```
25-env-and-logging/
├── main.go    # Main program with environment and logging examples
├── go.mod     # Module definition and dependencies
├── .env       # Sample environment variables file
├── logs/      # Directory for log files (created at runtime)
└── README.md  # Documentation
```

## Key Components

### 1. Environment Variables
- Loading from `.env` files using `godotenv`
- Default values for missing variables
- Type conversion (string, bool, int)
- Configuration validation
- Sensitive data handling
- Environment-specific settings

### 2. Standard Library Logging
- Basic logging with `log` package
- Log levels and formatting
- File-based logging
- Multi-writer setup (file + stdout)
- Error handling

### 3. Structured Logging (zerolog)
- JSON-structured logging
- Multiple log levels
- Context-rich logging
- Performance optimization
- Pretty printing for development

## Environment Variables

### Configuration Options
```env
# Logging configuration
LOG_LEVEL=debug|info|warn|error
LOG_FORMAT=json|pretty

# Application settings
PORT=3000
ENV=development|staging|production
ENABLE_FEATURE=true|false

# Sensitive information
API_KEY=your-secret-key
DATABASE_URL=your-database-url
```

### Best Practices

1. **Security**
   - Never commit sensitive values
   - Use environment-specific files
   - Mask secrets in logs
   - Validate required values

2. **Organization**
   - Group related variables
   - Use clear naming conventions
   - Document all options
   - Provide sensible defaults

3. **Type Safety**
   - Convert to appropriate types
   - Validate values early
   - Handle missing variables
   - Use strong typing

## Logging

### Log Levels
1. **Trace**: Very detailed information
2. **Debug**: Debugging information
3. **Info**: General information
4. **Warn**: Warning messages
5. **Error**: Error conditions
6. **Fatal**: Critical errors

### Best Practices

1. **Structure**
   - Use consistent formats
   - Include relevant context
   - Add timestamps
   - Include source information

2. **Performance**
   - Use appropriate log levels
   - Buffer log writes
   - Rotate log files
   - Clean up old logs

3. **Security**
   - Mask sensitive data
   - Control log access
   - Validate log data
   - Monitor log size

4. **Operations**
   - Use machine-readable format
   - Include correlation IDs
   - Add request tracking
   - Enable log aggregation

## Common Patterns

1. **Environment Loading**
   ```go
   if err := godotenv.Load(); err != nil {
       log.Printf("Warning: .env file not found")
   }
   ```

2. **Configuration Structure**
   ```go
   type Config struct {
       LogLevel  string
       APIKey    string
       Port      int
       Debug     bool
   }
   ```

3. **Structured Logging**
   ```go
   logger.Info().
       Str("user", "john").
       Int("age", 30).
       Msg("User logged in")
   ```

4. **Error Logging**
   ```go
   logger.Error().
       Err(err).
       Str("component", "database").
       Msg("Connection failed")
   ```

## Running the Example

1. **Setup**
   ```bash
   # Copy sample environment file
   cp .env.example .env

   # Install dependencies
   go mod tidy
   ```

2. **Run the Example**
   ```bash
   go run main.go
   ```

3. **Test Different Configurations**
   ```bash
   # Change log level
   LOG_LEVEL=debug go run main.go

   # Change log format
   LOG_FORMAT=pretty go run main.go

   # Set to production
   ENV=production go run main.go
   ```

## Additional Resources
- [godotenv Documentation](https://github.com/joho/godotenv)
- [zerolog Documentation](https://github.com/rs/zerolog)
- [12 Factor App - Config](https://12factor.net/config)
- [Go Logging Best Practices](https://www.honeybadger.io/blog/golang-logging/) 
