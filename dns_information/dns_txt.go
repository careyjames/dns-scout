package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
	"github.com/miekg/dns"
)

// GetTXT fetches the TXT records for a given domain.
func GetTXT(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeTXT, "8.8.8.8:53")
}

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
		fmt.Printf(color.Blue(" TXT Records: ") + constants.Newline)
		for _, record := range txt {
			isValidSPF := strings.HasPrefix(record, "v=spf1")
			coloredRecord := colorCodeSPFRecord(record, isValidSPF)
			fmt.Printf(" %s\n", coloredRecord)
		}
	} else {
		fmt.Printf("\033[38;5;39m TXT Records: \033[0m\033[38;5;88mNone\033[0m\n")
	}
}
