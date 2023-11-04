package dnsinformation

import (
	"fmt"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
	"github.com/miekg/dns"
)

// GetTXT fetches the TXT records for a given domain.
func GetTXT(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeTXT, constants.GooglePublicDNS)
}

// GetTXTFromAllOption fetches the TXT records for a given domain from ns lookup too.
func GetTXTFromAllOption(domain string) ([]string, error) {
	txtRecords, err := GetTXT(domain)
	if err != nil {
		return nil, err
	}
	if len(txtRecords) <= 0 {
		txtRecords, err = GetTXTRecordNSLookup(domain)
		if err != nil {
			return nil, err
		}
	}
	return txtRecords, nil
}

// GetTXTPrompt
func GetTXTPrompt(input string) {
	txt, _ := GetTXTFromAllOption(input)
	if len(txt) > 0 {
		fmt.Printf(color.Blue(" TXT   🟢: ") + constants.Newline)
		for _, record := range txt {
			coloredRecord := colorCodeSPFRecord(record)
			fmt.Printf(" %s\n", coloredRecord)
		}
	} else {
		fmt.Printf(color.Blue(" TXT   ❌: ") + color.Red("None") + constants.Newline)
	}
}
