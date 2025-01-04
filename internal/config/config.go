package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds the configuration for the application
type Config struct {
	Provider      string        // cloudflare or namecheap
	APIKey        string        // API Key for the chosen provider
	Email         string        // Email for Cloudflare (optional for Namecheap)
	Domain        string        // Domain to update
	Hostnames     []string      // List of hostnames to update
	CheckInterval time.Duration // Interval for checking IP changes
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Read provider
	provider := os.Getenv("PDYNIP_PROVIDER")
	if provider == "" {
		return nil, errors.New("PDYNIP_PROVIDER is required (e.g., 'cloudflare' or 'namecheap')")
	}

	// Read API key
	apiKey := os.Getenv("PDYNIP_API_KEY")
	if apiKey == "" {
		return nil, errors.New("PDYNIP_API_KEY is required")
	}

	// Read email (only needed for Cloudflare)
	email := os.Getenv("PDYNIP_EMAIL")

	// Read domain
	domain := os.Getenv("PDYNIP_DOMAIN")
	if domain == "" {
		return nil, errors.New("PDYNIP_DOMAIN is required")
	}

	// Read hostnames
	hostnames := os.Getenv("PDYNIP_HOSTNAMES")
	if hostnames == "" {
		return nil, errors.New("PDYNIP_HOSTNAMES is required (comma-separated list of hostnames)")
	}
	hostnameList := strings.Split(hostnames, ",")

	// Read check interval (optional, default 5 minutes)
	checkIntervalStr := os.Getenv("PDYNIP_CHECK_INTERVAL")
	checkInterval := 5 * time.Minute
	if checkIntervalStr != "" {
		interval, err := strconv.Atoi(checkIntervalStr)
		if err != nil || interval <= 0 {
			return nil, errors.New("PDYNIP_CHECK_INTERVAL must be a positive integer representing seconds")
		}
		checkInterval = time.Duration(interval) * time.Second
	}

	// Return the parsed configuration
	return &Config{
		Provider:      strings.ToLower(provider),
		APIKey:        apiKey,
		Email:         email,
		Domain:        domain,
		Hostnames:     hostnameList,
		CheckInterval: checkInterval,
	}, nil
}
