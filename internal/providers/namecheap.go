// Package providers implements DNS provider integrations for dynamic IP updates.
// This file contains the implementation for the Namecheap provider.
//
// NamecheapProvider interacts with Namecheap's Dynamic DNS API to update DNS records.
// It supports updating A records for a specified domain and hostname.

package providers

import (
	"encoding/xml"
	//"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// NamecheapProvider represents a Namecheap DNS provider
type NamecheapProvider struct {
	APIKey string // Dynamic DNS password from Namecheap
}

// namecheapAPIEndpoint is the base URL for Namecheap's Dynamic DNS API
const namecheapAPIEndpoint = "https://dynamicdns.park-your-domain.com/update"

// NewNamecheapProvider initializes a new NamecheapProvider
func NewNamecheapProvider(apiKey string) *NamecheapProvider {
	return &NamecheapProvider{
		APIKey: apiKey,
	}
}

// UpdateDNSRecord updates a DNS record for a given domain and hostname
func (nc *NamecheapProvider) UpdateDNSRecord(domain, hostname, ip string) error {
	// Build the request URL
	params := url.Values{}
	params.Set("host", hostname)
	params.Set("domain", domain)
	params.Set("password", nc.APIKey)
	params.Set("ip", ip)

	requestURL := fmt.Sprintf("%s?%s", namecheapAPIEndpoint, params.Encode())

	// Perform the HTTP request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(requestURL)
	if err != nil {
		return fmt.Errorf("failed to contact Namecheap API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 response code from Namecheap API: %d", resp.StatusCode)
	}

	// Parse the response for success/failure
	var result struct {
		ErrCount string `xml:"ErrCount"`
		Errors   []struct {
			Description string `xml:"Description"`
		} `xml:"errors>error"`
	}
	if err := parseXMLResponse(resp.Body, &result); err != nil {
		return fmt.Errorf("failed to parse Namecheap API response: %v", err)
	}

	if result.ErrCount != "0" {
		errMessages := []string{}
		for _, err := range result.Errors {
			errMessages = append(errMessages, err.Description)
		}
		return fmt.Errorf("Namecheap API returned errors: %s", errMessages)
	}

	return nil
}

// parseXMLResponse parses an XML response into a Go struct
func parseXMLResponse(body io.Reader, out interface{}) error {
	decoder := xml.NewDecoder(body)
	return decoder.Decode(out)
}
