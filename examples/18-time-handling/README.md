# Time Handling in Go

This example demonstrates comprehensive time handling in Go, including working with time values, epochs, formatting, parsing, and time zones.

## Key Components

### 1. Basic Time Operations
- Getting current time
- Accessing time components (year, month, day, etc.)
- Working with Unix timestamps
- Basic time zone conversions

### 2. Time Arithmetic
- Adding and subtracting durations
- Comparing times
- Rounding and truncating times
- Date calculations

### 3. Epoch Operations
- Working with Unix epochs
- Converting between time and epoch
- Handling nanosecond precision
- Important epoch dates (Y2K, 32-bit limit)

### 4. Time Formatting and Parsing
- Standard time formats (ANSIC, RFC3339, etc.)
- Custom format strings
- Parsing time strings
- Location-aware parsing

### 5. Time Zone Handling
- Working with different time zones
- Loading location data
- Handling DST transitions
- Time zone offsets

## Best Practices

1. **Time Storage**
   - Always store times in UTC internally
   - Convert to local time only for display
   - Use time.Time for timestamps, not string representations

2. **Format Strings**
   - Use predefined formats when possible (time.RFC3339, etc.)
   - Remember Go's reference time: "2006-01-02 15:04:05 -0700 MST"
   - Document custom format strings clearly

3. **Error Handling**
   - Always check errors when parsing times
   - Handle time zone loading errors gracefully
   - Validate input before parsing

## Memory Considerations

1. **Time.Time Structure**
   - time.Time is a struct containing both wall clock and monotonic clock
   - Efficient for most operations
   - Consider using Unix timestamps for storage optimization

2. **Location Data**
   - Time zone database is loaded into memory
   - Cache frequently used locations
   - Consider memory usage in location-heavy applications

## Common Patterns

1. **Time Comparison**
   ```go
   if timeA.Before(timeB) {
       // Handle earlier time
   }
   ```

2. **Duration Calculation**
   ```go
   duration := timeB.Sub(timeA)
   if duration > 24*time.Hour {
       // Handle day-long duration
   }
   ```

3. **Time Zone Conversion**
   ```go
   loc, _ := time.LoadLocation("America/New_York")
   nyTime := utcTime.In(loc)
   ```

## Safety Considerations

1. **Time Zone Safety**
   - Always validate time zone names
   - Handle missing time zone data gracefully
   - Be aware of DST transitions

2. **Parsing Safety**
   - Validate input format before parsing
   - Handle invalid dates appropriately
   - Consider time zone ambiguity during DST

## Testing Considerations

1. **Time-Based Testing**
   - Use fixed times for reproducible tests
   - Mock time.Now() for consistent results
   - Test edge cases (DST transitions, leap years)

2. **Time Zone Testing**
   - Test with multiple time zones
   - Include DST transition cases
   - Test invalid time zone handling

## Common Use Cases

1. **Event Scheduling**
   - Calculating future dates
   - Handling recurring events
   - Time zone aware scheduling

2. **Duration Tracking**
   - Measuring elapsed time
   - Computing time differences
   - Performance monitoring

3. **Log Analysis**
   - Parsing timestamp formats
   - Time-based filtering
   - Time zone normalization

## Advanced Topics

1. **Monotonic Clock**
   - Understanding wall vs monotonic time
   - Using time for duration measurements
   - Handling clock adjustments

2. **Custom Time Types**
   - Implementing time.Marshaler interface
   - Custom formatting methods
   - Time range types

3. **Performance Optimization**
   - Caching time zone data
   - Efficient time comparisons
   - Batch time operations

## Running the Program

```bash
go run main.go
```

The program will demonstrate various time handling operations, including:
- Basic time operations
- Time arithmetic
- Epoch handling
- Time formatting and parsing
- Time zone management

## Further Reading

- [Go time package documentation](https://golang.org/pkg/time/)
- [IANA Time Zone Database](https://www.iana.org/time-zones)
- [RFC 3339 - Date and Time on the Internet](https://tools.ietf.org/html/rfc3339) 
