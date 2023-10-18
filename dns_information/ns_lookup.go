package dnsinformation

import (
	"net"
)

// GetTXTRecordNSLookup fetch txt ns lookup
func GetTXTRecordNSLookup(domain string) ([]string, error) {
	dns_txt_records, err := net.LookupTXT(domain)
	records := []string{}
	if err == nil && len(dns_txt_records) > 0 {
		for _, record := range dns_txt_records {
			records = append(records, record)
		}
	} else {
		return records, err
	}
	return records, nil
}

// GetTXTRecordNSLookup fetch txt ns lookup
func GetDMARCRecordNSLookup(domain string) ([]string, error) {
	dns_txt_records, err := net.LookupTXT("_dmarc." + domain)
	records := []string{}
	if err == nil && len(dns_txt_records) > 0 {
		for _, record := range dns_txt_records {
			records = append(records, record)
		}
	} else {
		return records, err
	}
	return records, nil
}
