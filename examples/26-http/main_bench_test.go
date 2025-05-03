package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHealthHandler(b *testing.B) {
	// Create handler
	handler := http.HandlerFunc(handleHealth)

	// Reset timer before the loop
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req := httptest.NewRequest("GET", "/health", nil)
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkCreateUser(b *testing.B) {
	// Initialize user store
	userStore = make(map[int]User)
	nextID = 1

	// Create handler
	handler := http.HandlerFunc(handleUsers)

	// Prepare user data
	user := User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	userData, _ := json.Marshal(user)

	// Reset timer before the loop
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(userData))
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkGetUsers(b *testing.B) {
	// Initialize user store with some data
	userStore = make(map[int]User)
	for i := 1; i <= 100; i++ {
		userStore[i] = User{
			ID:    i,
			Name:  "Test User",
			Email: "test@example.com",
		}
	}

	// Create handler
	handler := http.HandlerFunc(handleUsers)

	// Reset timer before the loop
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req := httptest.NewRequest("GET", "/api/users", nil)
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkGetSingleUser(b *testing.B) {
	// Initialize user store with test data
	userStore = make(map[int]User)
	userStore[1] = User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	// Create handler
	handler := http.HandlerFunc(handleUsers)

	// Reset timer before the loop
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req := httptest.NewRequest("GET", "/api/users/1", nil)
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkUpdateUser(b *testing.B) {
	// Initialize user store with test data
	userStore = make(map[int]User)
	userStore[1] = User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	// Create handler
	handler := http.HandlerFunc(handleUsers)

	// Prepare update data
	updateUser := User{
		Name:  "Updated User",
		Email: "updated@example.com",
	}
	userData, _ := json.Marshal(updateUser)

	// Reset timer before the loop
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Create request
		req := httptest.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(userData))
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}

func BenchmarkDeleteUser(b *testing.B) {
	// Create handler
	handler := http.HandlerFunc(handleUsers)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Reset user store for each iteration
		b.StopTimer()
		userStore = make(map[int]User)
		userStore[1] = User{
			ID:    1,
			Name:  "Test User",
			Email: "test@example.com",
		}
		b.StartTimer()

		// Create request
		req := httptest.NewRequest("DELETE", "/api/users/1", nil)
		rr := httptest.NewRecorder()

		// Call handler
		handler.ServeHTTP(rr, req)
	}
}
