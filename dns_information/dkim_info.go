package dnsinformation

import (
	"fmt"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

const (
	DKIMValid = "v=DKIM1"
)

var (
	DKIMSelectors = []string{
		"default._domainkey",
		"google._domainkey",
		"mail._domainkey",
		"selector1._domainkey",
		"selector2._domainkey",
	}
)

// getDKIM fetches the DMARC record for a given domain.
func getDKIM(domain string, selector string) (string, error) {
	txtRecords, _ := GetTXTRecordNSLookup(selector + "." + domain)
	for _, record := range txtRecords {
		if hasDKIMRecord(record) {
			return record, nil
		}
	}
	return "", nil
}

// GetDKIMPrompt fetches the DKIM record for a given domain.
func GetDKIMPrompt(input string) {
	flag := false
	dkimPrompt := ""
	for index, selector := range DKIMSelectors {
		dkim, _ := getDKIM(input, selector)
		if dkim != "" {
			if !flag {
				dkimPrompt = ""
			}
			flag = true
			if isValidDKIM(dkim) {
				formattedDMARC := formatLongText(dkim, 80, " ")
				if index > 0 && len(dkimPrompt) > 0 {
					dkimPrompt += " " + color.Grey(selector+".") + color.Grey(formattedDMARC) + constants.Newline
				} else {
					dkimPrompt += " " + color.Grey(selector+".") + color.Grey(formattedDMARC) + constants.Newline
				}
			} else {
				dkimPrompt += " " + color.Grey(selector+".") + color.Red(dkim[7:]) + constants.Newline
			}
		} else {
			if !flag {
				dkimPrompt = color.Red(" None") + constants.Newline
			}
		}
	}
	if flag {
		fmt.Printf(color.Blue(" DKIM  ✅:") + dkimPrompt)
	} else {
		fmt.Printf(color.Blue(" DKIM  ❌:") + dkimPrompt)
	}
}

func hasDKIMRecord(record string) bool {
	return (len(record) > 7 && record[:7] == DKIMValid) || (len(record) > 5 && record[:5] == "vDKIM")
}

func isValidDKIM(record string) bool {
	if len(record) > 7 && record[:7] == DKIMValid {
		return true
	}
	return false
}
