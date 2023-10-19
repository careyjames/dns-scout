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

// getDMKI fetches the DMARC record for a given domain.
func getDKIM(domain string) (string, error) {
	var txtRecords []string
	for _, selector := range DKIMSelectors {
		txtRecord, _ := GetTXTRecordNSLookup(selector + "." + domain)
		txtRecords = append(txtRecords, txtRecord...)
	}
	for _, record := range txtRecords {
		if hasDKIMRecord(record) {
			return record, nil
		}
	}
	return "", nil
}

// GetDKIMPrompt fetches the DKIM record for a given domain.
func GetDKIMPrompt(input string) {
	dkmi, _ := getDKIM(input)
	if dkmi != "" {
		if isValidDKIM(dkmi) {
			formattedDMARC := formatLongText(dkmi, 80, " ")
			fmt.Printf(color.Blue(" DKIM Record: ") + color.Green(formattedDMARC) + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" DKIM Record: ") + color.Green(dkmi[7:]) + constants.Newline)
		}
	} else {
		fmt.Printf(color.Blue(" DKIM Record: ") + color.Red("None") + constants.Newline)
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
