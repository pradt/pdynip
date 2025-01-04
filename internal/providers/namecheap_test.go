package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNamecheapUpdateDNSRecord(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<?xml version="1.0"?><interface-response><ErrCount>0</ErrCount></interface-response>`))
	}))
	defer mockServer.Close()

	namecheapAPIEndpoint = mockServer.URL

	provider := NewNamecheapProvider("test-api-key")
	err := provider.UpdateDNSRecord("example.com", "www", "203.0.113.1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
