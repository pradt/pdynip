package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDetectPublicIP(t *testing.T) {
	// Mock the IP service
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("203.0.113.1"))
	}))
	defer mockServer.Close()

	// Replace the service URL with the mock server
	ipServiceURL = mockServer.URL

	ip, err := DetectPublicIP()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ip != "203.0.113.1" {
		t.Errorf("Expected IP '203.0.113.1', got '%s'", ip)
	}
}
