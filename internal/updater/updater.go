// Package updater coordinates the DNS record updates for dynamic IP changes.
// It works with any provider that implements the Provider interface.

package updater

import (
	"log"
	"pdynip/internal/providers"
)

// UpdateDNS performs a DNS update using the specified provider
func UpdateDNS(provider providers.Provider, domain, hostname, ip string) error {
	log.Printf("Starting update for hostname: %s in domain: %s with IP: %s", hostname, domain, ip)

	// Attempt to update the DNS record
	err := provider.UpdateDNSRecord(domain, hostname, ip)
	if err != nil {
		log.Printf("Failed to update DNS record for hostname: %s in domain: %s - %v", hostname, domain, err)
		return err
	}

	log.Printf("Successfully updated DNS record for hostname: %s in domain: %s with IP: %s", hostname, domain, ip)
	return nil
}
