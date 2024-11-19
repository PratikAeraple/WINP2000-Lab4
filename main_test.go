package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileServer(t *testing.T) {
	// Setup: Create the `templates` directory and add a test file
	staticDir := "./templates"
	err := os.MkdirAll(staticDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create static directory: %v", err)
	}
	defer os.RemoveAll(staticDir) // Clean up after the test

	// Add a sample file to the directory
	fileName := "index.html"
	fileContent := "<html><body><h1>Test Page</h1></body></html>"
	err = os.WriteFile(staticDir+"/"+fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Set up the file server and start a test server
	fs := http.FileServer(http.Dir(staticDir))
	ts := httptest.NewServer(fs)
	defer ts.Close()

	// Perform a GET request for the existing file
	resp, err := http.Get(ts.URL + "/index.html")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got: %d", resp.StatusCode)
	}

	// Perform a GET request for a non-existent file
	resp, err = http.Get(ts.URL + "/missing.html")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status for a non-existent file
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status 404 Not Found, got: %d", resp.StatusCode)
	}
}
