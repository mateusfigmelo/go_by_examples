package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Global variables
var (
	userStore = make(map[int]User)
	nextID    = 1
	mu        sync.RWMutex
)

// HTTP Server Examples

// setupServer creates and configures the HTTP server
func setupServer() *http.Server {
	// Create a new mux for routing
	mux := http.NewServeMux()

	// Basic handlers
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/health", handleHealth)

	// RESTful API endpoints
	mux.HandleFunc("/api/users", handleUsers)
	mux.HandleFunc("/api/users/", handleUsers)

	// Create server with configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      logMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}

// Middleware for logging
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// handleHome handles the root endpoint
func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the Go HTTP Example!")
}

// handleHealth handles the health check endpoint
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

// handleUsers handles CRUD operations for users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from path if present
	var userID int
	var err error

	// Split path and get last segment if it exists
	parts := strings.Split(strings.TrimRight(r.URL.Path, "/"), "/")
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if lastPart != "users" {
			userID, err = strconv.Atoi(lastPart)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}
		}
	}

	switch r.Method {
	case http.MethodGet:
		if userID > 0 {
			// Get single user
			mu.RLock()
			user, exists := userStore[userID]
			mu.RUnlock()
			if !exists {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)
		} else {
			// Get all users
			mu.RLock()
			users := make([]User, 0, len(userStore))
			for _, user := range userStore {
				users = append(users, user)
			}
			mu.RUnlock()
			json.NewEncoder(w).Encode(users)
		}

	case http.MethodPost:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		user.ID = nextID
		userStore[nextID] = user
		nextID++
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	case http.MethodPut:
		if userID == 0 {
			http.Error(w, "User ID required", http.StatusBadRequest)
			return
		}

		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		if _, exists := userStore[userID]; !exists {
			mu.Unlock()
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		user.ID = userID
		userStore[userID] = user
		mu.Unlock()

		json.NewEncoder(w).Encode(user)

	case http.MethodDelete:
		if userID == 0 {
			http.Error(w, "User ID required", http.StatusBadRequest)
			return
		}

		mu.Lock()
		if _, exists := userStore[userID]; !exists {
			mu.Unlock()
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		delete(userStore, userID)
		mu.Unlock()

		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HTTP Client Examples

// setupClient creates and configures the HTTP client
func setupClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
}

// demonstrateClientOperations shows various HTTP client operations
func demonstrateClientOperations(baseURL string) {
	client := setupClient()

	// GET request
	resp, err := client.Get(baseURL + "/health")
	if err != nil {
		log.Printf("GET request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return
	}
	log.Printf("Health check response: %v", result)

	// POST request
	user := User{Name: "John Doe", Email: "john@example.com"}
	userJSON, _ := json.Marshal(user)
	resp, err = client.Post(baseURL+"/api/users", "application/json", strings.NewReader(string(userJSON)))
	if err != nil {
		log.Printf("POST request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	var createdUser User
	if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return
	}
	log.Printf("Created user: %+v", createdUser)
}

func main() {
	fmt.Println("=== HTTP Server and Client Example ===")

	// Start server
	server := setupServer()
	go func() {
		fmt.Printf("\nServer starting on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	// Wait for server to start
	time.Sleep(time.Second)

	// Demo client operations
	baseURL := "http://localhost:8080"
	demonstrateClientOperations(baseURL)

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	fmt.Println("\nShutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}
}
