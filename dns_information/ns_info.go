package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
	"github.com/miekg/dns"
)

// GetNS fetches the NS records for a given domain.
func GetNS(domain string) ([]string, error) {
	googleRecords, err1 := QueryDNS(domain, dns.TypeNS, constants.GooglePublicDNS)
	cloudflareRecords, err2 := QueryDNS(domain, dns.TypeNS, constants.CloudflarePublicDNS)

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
		fmt.Printf(color.Blue(" NS    🟢: ") + color.Grey(strings.Join(ns, ", ")) + constants.Newline)
	} else {
		fmt.Printf(color.Blue(" NS    ❌: ") + color.Red("None") + constants.Newline)
	}
}
