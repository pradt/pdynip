package ip

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const ipServiceURL = "https://api.ipify.org" // Service to fetch the public IP

// DetectPublicIP fetches the current public IP address
func DetectPublicIP() (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Set timeout to avoid hanging requests
	}

	resp, err := client.Get(ipServiceURL)
	if err != nil {
		return "", errors.New("failed to fetch public IP: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch public IP: received non-200 response code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body: " + err.Error())
	}

	ip := strings.TrimSpace(string(body))
	if ip == "" {
		return "", errors.New("failed to detect public IP: empty response")
	}

	return ip, nil
}
