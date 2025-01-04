package updater

import (
	"errors"
	"testing"
)

type MockProvider struct {
	ShouldFail bool
}

func (m *MockProvider) UpdateDNSRecord(domain, hostname, ip string) error {
	if m.ShouldFail {
		return errors.New("mock provider failure")
	}
	return nil
}

func TestUpdateDNS(t *testing.T) {
	mockProvider := &MockProvider{ShouldFail: false}

	err := UpdateDNS(mockProvider, "example.com", "www", "203.0.113.1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	mockProvider.ShouldFail = true
	err = UpdateDNS(mockProvider, "example.com", "www", "203.0.113.1")
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}
}
