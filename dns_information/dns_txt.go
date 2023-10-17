package dnsinformation

import "github.com/miekg/dns"

// GetTXT fetches the TXT records for a given domain.
func GetTXT(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeTXT, "8.8.8.8:53")
}
