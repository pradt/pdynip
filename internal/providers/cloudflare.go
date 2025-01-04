// Package providers implements DNS provider integrations for dynamic IP updates.
// This file contains the implementation for the Cloudflare provider.
//
// CloudflareProvider interacts with the Cloudflare API to update DNS records.
// It supports updating A records for a specified domain and hostname.

package providers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// CloudflareProvider represents a Cloudflare DNS provider
type CloudflareProvider struct {
	APIKey string
	Email  string
}

// cloudflareAPIEndpoint is the base URL for Cloudflare's API
var cloudflareAPIEndpoint = "https://api.cloudflare.com/client/v4"

// NewCloudflareProvider initializes a new CloudflareProvider
func NewCloudflareProvider(apiKey, email string) *CloudflareProvider {
	return &CloudflareProvider{
		APIKey: apiKey,
		Email:  email,
	}
}

// UpdateDNSRecord updates a DNS record for a given domain and hostname
func (cf *CloudflareProvider) UpdateDNSRecord(domain, hostname, ip string) error {
	zoneID, err := cf.getZoneID(domain)
	if err != nil {
		return fmt.Errorf("failed to get zone ID: %v", err)
	}

	recordID, err := cf.getDNSRecordID(zoneID, hostname)
	if err != nil {
		return fmt.Errorf("failed to get DNS record ID: %v", err)
	}

	if err := cf.updateDNSRecord(zoneID, recordID, hostname, ip); err != nil {
		return fmt.Errorf("failed to update DNS record: %v", err)
	}

	return nil
}

// getZoneID retrieves the Zone ID for the given domain
func (cf *CloudflareProvider) getZoneID(domain string) (string, error) {
	url := fmt.Sprintf("%s/zones?name=%s", cloudflareAPIEndpoint, domain)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+cf.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

	var result struct {
		Success bool `json:"success"`
		Errors  []struct {
			Message string `json:"message"`
		} `json:"errors"`
		Result []struct {
			ID string `json:"id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if !result.Success || len(result.Result) == 0 {
		return "", errors.New("failed to retrieve zone ID")
	}

	return result.Result[0].ID, nil
}

// getDNSRecordID retrieves the DNS record ID for the given zone and hostname
func (cf *CloudflareProvider) getDNSRecordID(zoneID, hostname string) (string, error) {
	url := fmt.Sprintf("%s/zones/%s/dns_records?name=%s", cloudflareAPIEndpoint, zoneID, hostname)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+cf.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

	var result struct {
		Success bool `json:"success"`
		Errors  []struct {
			Message string `json:"message"`
		} `json:"errors"`
		Result []struct {
			ID string `json:"id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if !result.Success || len(result.Result) == 0 {
		return "", errors.New("failed to retrieve DNS record ID")
	}

	return result.Result[0].ID, nil
}

// updateDNSRecord updates the DNS record with the new IP
func (cf *CloudflareProvider) updateDNSRecord(zoneID, recordID, hostname, ip string) error {
	url := fmt.Sprintf("%s/zones/%s/dns_records/%s", cloudflareAPIEndpoint, zoneID, recordID)

	payload := map[string]interface{}{
		"type":    "A",
		"name":    hostname,
		"content": ip,
		"ttl":     1,
		"proxied": false,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+cf.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 response code: %d", resp.StatusCode)
	}

	var result struct {
		Success bool `json:"success"`
		Errors  []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Success {
		return errors.New("failed to update DNS record")
	}

	return nil
}
