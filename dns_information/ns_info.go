package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

// GetNS fetches the NS records for a given domain.
func GetNS(domain string) ([]string, error) {
	googleRecords, err1 := QueryDNS(domain, dns.TypeNS, "8.8.8.8:53")
	cloudflareRecords, err2 := QueryDNS(domain, dns.TypeNS, "1.1.1.1:53")

	if err1 != nil && err2 != nil {
		return nil, fmt.Errorf("both DNS queries failed")
	}

	// Merge and deduplicate records
	recordMap := make(map[string]bool)
	for _, record := range googleRecords {
		recordMap[record] = true
	}
	for _, record := range cloudflareRecords {
		recordMap[record] = true
	}

	var mergedRecords []string
	for record := range recordMap {
		mergedRecords = append(mergedRecords, record)
	}

	return mergedRecords, nil
}

// GetNSPrompt is ns prompt for response
func GetNSPrompt(input string) {
	ns, _ := GetNS(input)
	if len(ns) > 0 {
		fmt.Printf("\033[38;5;39m Name Servers: \033[38;5;78m%s\033[0m\n", strings.Join(ns, ", "))
	} else {
		fmt.Printf("\033[38;5;39m Name Servers: \033[0m\033[38;5;88mNone\033[0m\n")
	}
}
