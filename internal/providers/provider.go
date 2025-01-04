package providers

type Provider interface {
	UpdateDNSRecord(domain, hostname, ip string) error
}
