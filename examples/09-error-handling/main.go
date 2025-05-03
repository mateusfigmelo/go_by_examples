package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Custom error types
type ValidationError struct {
	Field string
	Issue string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Issue)
}

// NotFoundError represents a not found error
type NotFoundError struct {
	Type string
	ID   string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Type, e.ID)
}

// Custom error with additional context
type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("query '%s' failed: %v", e.Query, e.Err)
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

// Domain-specific errors
var (
	ErrInvalidInput    = errors.New("invalid input provided")
	ErrNotAuthorized   = errors.New("not authorized to perform this action")
	ErrResourceExpired = errors.New("resource has expired")
)

// User represents a user in the system
type User struct {
	ID       string
	Name     string
	Age      int
	Email    string
	Password string
}

// ValidateUser demonstrates validation with custom errors
func ValidateUser(user *User) error {
	if user.Name == "" {
		return &ValidationError{
			Field: "name",
			Issue: "cannot be empty",
		}
	}

	if user.Age < 0 || user.Age > 150 {
		return &ValidationError{
			Field: "age",
			Issue: "must be between 0 and 150",
		}
	}

	if !strings.Contains(user.Email, "@") {
		return &ValidationError{
			Field: "email",
			Issue: "invalid email format",
		}
	}

	if len(user.Password) < 8 {
		return &ValidationError{
			Field: "password",
			Issue: "must be at least 8 characters",
		}
	}

	return nil
}

// FindUser demonstrates error wrapping and handling not found errors
func FindUser(id string) (*User, error) {
	// Simulate database lookup
	if id == "" {
		return nil, ErrInvalidInput
	}

	if id == "404" {
		return nil, &NotFoundError{
			Type: "User",
			ID:   id,
		}
	}

	// Simulate other errors
	if id == "expired" {
		return nil, fmt.Errorf("finding user: %w", ErrResourceExpired)
	}

	// Success case
	return &User{ID: id, Name: "John Doe"}, nil
}

// ExecuteQuery demonstrates error wrapping with context
func ExecuteQuery(query string) error {
	// Simulate database query
	if query == "" {
		return &QueryError{
			Query: query,
			Err:   ErrInvalidInput,
		}
	}

	if strings.Contains(query, "DROP") {
		return &QueryError{
			Query: query,
			Err:   ErrNotAuthorized,
		}
	}

	return nil
}

// ProcessFile demonstrates handling multiple error types
func ProcessFile(filename string) error {
	// Try to open file
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %w", err)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied: %w", err)
		}
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Try to read file
	_, err = io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// SafeOperation demonstrates panic recovery
func SafeOperation(data interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	// Simulate some risky operation
	if data == nil {
		panic("nil data provided")
	}

	return nil
}

func main() {
	fmt.Println("=== Go Error Handling Examples ===")

	// 1. Basic Error Handling
	fmt.Println("\n1. Basic Error Handling:")
	user := &User{
		Name:     "",
		Age:      200,
		Email:    "invalid-email",
		Password: "123",
	}

	if err := ValidateUser(user); err != nil {
		fmt.Printf("    Validation error: %v\n", err)
	}

	// 2. Custom Error Types
	fmt.Println("\n2. Custom Error Types:")
	user.Name = "John"
	if err := ValidateUser(user); err != nil {
		if validErr, ok := err.(*ValidationError); ok {
			fmt.Printf("    Field: %s, Issue: %s\n", validErr.Field, validErr.Issue)
		}
	}

	// 3. Error Wrapping and Unwrapping
	fmt.Println("\n3. Error Wrapping and Unwrapping:")
	_, err := FindUser("404")
	if err != nil {
		if notFound, ok := err.(*NotFoundError); ok {
			fmt.Printf("    Not found error: %v\n", notFound)
		}
	}

	_, err = FindUser("expired")
	if err != nil {
		if errors.Is(err, ErrResourceExpired) {
			fmt.Printf("    Resource expired error: %v\n", err)
		}
	}

	// 4. Error Context
	fmt.Println("\n4. Error Context:")
	err = ExecuteQuery("DROP TABLE users")
	if err != nil {
		if queryErr, ok := err.(*QueryError); ok {
			fmt.Printf("    Query error: %v\n", queryErr)
			if errors.Is(queryErr.Err, ErrNotAuthorized) {
				fmt.Printf("    Unauthorized operation detected\n")
			}
		}
	}

	// 5. File Operation Errors
	fmt.Println("\n5. File Operation Errors:")
	err = ProcessFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("    File error: %v\n", err)
	}

	// 6. Panic Recovery
	fmt.Println("\n6. Panic Recovery:")
	err = SafeOperation(nil)
	if err != nil {
		fmt.Printf("    Recovered from panic: %v\n", err)
	}

	// 7. Multiple Error Checks
	fmt.Println("\n7. Multiple Error Checks:")
	users := []string{"", "404", "expired", "valid"}
	for _, id := range users {
		user, err := FindUser(id)
		if err != nil {
			switch {
			case errors.Is(err, ErrInvalidInput):
				fmt.Printf("    Invalid input for ID: %s\n", id)
			case errors.Is(err, ErrResourceExpired):
				fmt.Printf("    Resource expired: %s\n", id)
			default:
				if notFound, ok := err.(*NotFoundError); ok {
					fmt.Printf("    %v\n", notFound)
				} else {
					fmt.Printf("    Unknown error: %v\n", err)
				}
			}
			continue
		}
		fmt.Printf("    Found user: %s\n", user.Name)
	}
}
