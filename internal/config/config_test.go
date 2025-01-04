package config

import (
	"os"
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("PDYNIP_PROVIDER", "cloudflare")
	os.Setenv("PDYNIP_API_KEY", "test-api-key")
	os.Setenv("PDYNIP_EMAIL", "test@example.com")
	os.Setenv("PDYNIP_DOMAIN", "example.com")
	os.Setenv("PDYNIP_HOSTNAMES", "www,api")
	os.Setenv("PDYNIP_CHECK_INTERVAL", "300")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.Provider != "cloudflare" {
		t.Errorf("Expected provider 'cloudflare', got '%s'", cfg.Provider)
	}

	if cfg.CheckInterval != 300*time.Second {
		t.Errorf("Expected check interval of 300 seconds, got %v", cfg.CheckInterval)
	}
}
