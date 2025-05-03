# HTTP Client and Server in Go

This example demonstrates how to create and use HTTP clients and servers in Go, including RESTful APIs, middleware, and various HTTP operations.

## Project Structure
```
26-http/
├── main.go    # Main program with HTTP client and server examples
└── README.md  # Documentation
```

## Key Components

### 1. HTTP Server
- Basic HTTP server setup
- Request routing with `http.ServeMux`
- Middleware implementation
- RESTful API endpoints
- JSON handling
- Error handling
- Graceful shutdown

### 2. HTTP Client
- Client configuration
- Various request types (GET, POST, PUT, DELETE)
- Custom headers
- Context usage
- Response handling
- Connection pooling

### 3. RESTful API
- Resource-based routing
- CRUD operations
- JSON request/response
- Status codes
- Error responses

## Server Features

### 1. Basic Setup
```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      handler,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  15 * time.Second,
}
```

### 2. Routing
```go
mux := http.NewServeMux()
mux.HandleFunc("/", handleHome)
mux.HandleFunc("/api/users", handleUsers)
```

### 3. Middleware
```go
func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Pre-processing
        next.ServeHTTP(w, r)
        // Post-processing
    })
}
```

### 4. Request Handling
```go
func handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Handle GET
    case http.MethodPost:
        // Handle POST
    }
}
```

## Client Features

### 1. Client Setup
```go
client := &http.Client{
    Timeout: time.Second * 10,
    Transport: &http.Transport{
        MaxIdleConns: 10,
        IdleConnTimeout: 30 * time.Second,
    },
}
```

### 2. Making Requests
```go
// GET request
resp, err := client.Get(url)

// POST request
resp, err := client.Post(url, contentType, body)

// Custom request
req, _ := http.NewRequestWithContext(ctx, method, url, body)
resp, err := client.Do(req)
```

## Best Practices

### 1. Server
- Use timeouts
- Implement graceful shutdown
- Handle all errors
- Use appropriate status codes
- Validate input
- Log requests
- Use middleware for cross-cutting concerns

### 2. Client
- Reuse clients
- Set timeouts
- Close response bodies
- Use context for cancellation
- Handle redirects
- Pool connections
- Handle errors appropriately

### 3. General
- Use HTTPS in production
- Validate content types
- Handle character encoding
- Follow REST conventions
- Use proper error formats
- Document API endpoints

## Common Patterns

### 1. JSON API
```go
// Request handling
var data struct {
    Name string `json:"name"`
}
json.NewDecoder(r.Body).Decode(&data)

// Response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
```

### 2. Error Handling
```go
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

### 3. Middleware Chain
```go
handler = logging(authentication(compression(handler)))
```

### 4. Context Usage
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
req = req.WithContext(ctx)
```

## Running the Example

1. **Start the Server**
   ```bash
   go run main.go
   ```

2. **Test Endpoints**
   ```bash
   # Health check
   curl http://localhost:8080/health

   # Create user
   curl -X POST http://localhost:8080/api/users \
        -H "Content-Type: application/json" \
        -d '{"name":"John","email":"john@example.com"}'

   # Get users
   curl http://localhost:8080/api/users

   # Update user
   curl -X PUT http://localhost:8080/api/users/1 \
        -H "Content-Type: application/json" \
        -d '{"name":"John Updated","email":"john@example.com"}'

   # Delete user
   curl -X DELETE http://localhost:8080/api/users/1
   ```

## Additional Resources
- [net/http Package Documentation](https://golang.org/pkg/net/http/)
- [REST API Design Best Practices](https://golang.org/doc/articles/wiki/)
- [HTTP/2 in Go](https://golang.org/doc/articles/h2c)
- [Context Package Documentation](https://golang.org/pkg/context/) 
