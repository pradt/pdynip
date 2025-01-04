package main

import (
	"log"
	"strings"
	"time"

	"pdynip/internal/config"
	"pdynip/internal/ip"
	"pdynip/internal/providers"
	"pdynip/internal/updater"
)

func main() {
	// Load configuration from environment variables or CLI arguments
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize provider based on configuration
	var provider providers.Provider
	switch strings.ToLower(cfg.Provider) {
	case "cloudflare":
		provider = providers.NewCloudflareProvider(cfg.APIKey, cfg.Email)
	case "namecheap":
		provider = providers.NewNamecheapProvider(cfg.APIKey)
	default:
		log.Fatalf("Unsupported provider: %s", cfg.Provider)
	}

	// Detect initial public IP
	log.Println("Detecting current public IP...")
	currentIP, err := ip.DetectPublicIP()
	if err != nil {
		log.Fatalf("Error detecting public IP: %v", err)
	}
	log.Printf("Detected initial public IP: %s", currentIP)

	// Start the IP monitoring loop
	for {
		newIP, err := ip.DetectPublicIP()
		if err != nil {
			log.Printf("Error detecting public IP: %v", err)
			time.Sleep(cfg.CheckInterval)
			continue
		}

		if newIP != currentIP {
			log.Printf("Detected IP change from %s to %s", currentIP, newIP)

			// Update all hostnames for the provider
			for _, hostname := range cfg.Hostnames {
				log.Printf("Updating %s hostname %s with IP %s", cfg.Provider, hostname, newIP)
				err = updater.UpdateDNS(provider, cfg.Domain, hostname, newIP)
				if err != nil {
					log.Printf("Error updating hostname %s: %v", hostname, err)
				} else {
					log.Printf("Successfully updated hostname %s to IP %s", hostname, newIP)
				}
			}

			// Update the current IP
			currentIP = newIP
		}

		// Wait for the next check
		time.Sleep(cfg.CheckInterval)
	}
}
