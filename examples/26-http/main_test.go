package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHealth)

	// Call the handler directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body (trim newlines for comparison)
	expected := `{"status":"healthy"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserHandlers(t *testing.T) {
	// Initialize user store
	userStore = make(map[int]User)
	nextID = 1

	tests := []struct {
		name           string
		method         string
		path           string
		body           interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Create User",
			method:         "POST",
			path:           "/api/users",
			body:           User{Name: "Test User", Email: "test@example.com"},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"id":1,"name":"Test User","email":"test@example.com"}`,
		},
		{
			name:           "Get Users",
			method:         "GET",
			path:           "/api/users",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Test User","email":"test@example.com"}]`,
		},
		{
			name:           "Get Single User",
			method:         "GET",
			path:           "/api/users/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Test User","email":"test@example.com"}`,
		},
		{
			name:           "Update User",
			method:         "PUT",
			path:           "/api/users/1",
			body:           User{Name: "Updated User", Email: "updated@example.com"},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Updated User","email":"updated@example.com"}`,
		},
		{
			name:           "Delete User",
			method:         "DELETE",
			path:           "/api/users/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User deleted successfully"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.body != nil {
				body, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatal(err)
				}
			}

			req, err := http.NewRequest(tt.method, tt.path, bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			// Set content type for requests with body
			if tt.body != nil {
				req.Header.Set("Content-Type", "application/json")
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handleUsers)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Compare JSON responses (trim newlines)
			if strings.TrimSpace(rr.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestClientOperations(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/test":
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"test response"}`))
		case "/error":
			http.Error(w, "test error", http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	// Test GET request
	t.Run("GET Request", func(t *testing.T) {
		client := &http.Client{}
		resp, err := client.Get(ts.URL + "/test")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", resp.Status)
		}

		var result map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}

		if result["message"] != "test response" {
			t.Errorf("expected message 'test response'; got %v", result["message"])
		}
	})

	// Test error handling
	t.Run("Error Response", func(t *testing.T) {
		client := &http.Client{}
		resp, err := client.Get(ts.URL + "/error")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status 500; got %v", resp.Status)
		}
	})
}

func ExamplehandleHealth() {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	handleHealth(w, req)
	fmt.Print(strings.TrimSpace(w.Body.String()))
	// Output: {"status":"healthy"}
}
