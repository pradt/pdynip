package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCloudflareUpdateDNSRecord(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true}`))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer mockServer.Close()

	cloudflareAPIEndpoint = mockServer.URL

	provider := NewCloudflareProvider("test-api-key", "test@example.com")
	err := provider.UpdateDNSRecord("example.com", "www", "203.0.113.1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
