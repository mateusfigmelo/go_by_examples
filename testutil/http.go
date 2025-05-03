package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// HTTPTestServer represents a test HTTP server
type HTTPTestServer struct {
	*httptest.Server
	LastRequest *http.Request
}

// NewHTTPTestServer creates a new test HTTP server with the given handler
func NewHTTPTestServer(handler http.HandlerFunc) *HTTPTestServer {
	ts := &HTTPTestServer{}
	ts.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts.LastRequest = r
		handler(w, r)
	}))
	return ts
}

// Get performs a GET request to the test server
func (ts *HTTPTestServer) Get(path string) (int, string, error) {
	resp, err := http.Get(ts.URL + path)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}

	return resp.StatusCode, string(body), nil
}
