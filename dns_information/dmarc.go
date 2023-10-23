package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

const (
	DMARCLookupString = "_dmarc."
	DMARCValid        = "v=DMARC1"
)

// getDMARC fetches the DMARC record for a given domain.
func getDMARC(domain string) (string, error) {
	txtRecords, err := GetTXT(DMARCLookupString + domain)
	if len(txtRecords) <= 0 {
		txtRecords, _ = GetDMARCRecordNSLookup(domain)
	}
	if err != nil {
		return "", err
	}
	for _, record := range txtRecords {
		if HasDMARCRecord(record) {
			return record, nil
		}
	}
	return "", nil
}

func HasDMARCRecord(record string) bool {
	return (len(record) > 8 && record[:8] == DMARCValid) || (len(record) > 6 && record[:6] == "vDMARC")
}

// GetDMARCPrompt fetches the DMARC record for a given domain.
func GetDMARCPrompt(input string) {
	dmarc, _ := getDMARC(input)
	if dmarc != "" {
		if isValidDMARC(dmarc) {
			formattedDMARC := formatLongText(dmarc, 80, " ")
			fmt.Printf(color.Blue(" DMARC ✅:") + color.Grey(" ") + color.Grey(formattedDMARC) + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" DMARC ❌:") + color.Red(" ") + color.Red(dmarc[8:]) + constants.Newline)
		}
	} else {
		fmt.Printf(color.Blue(" DMARC ❌:") + color.Red(" None") + constants.Newline)
	}
}

func isValidDMARC(record string) bool {
	if len(record) > 8 && record[:8] == DMARCValid {
		return true
	}
	return false
}

// formatLongText formats long text strings for better readability.
func formatLongText(text string, threshold int, indent string) string {
	if len(text) <= threshold {
		return text
	}

	var result strings.Builder
	for len(text) > threshold {
		splitAt := strings.LastIndex(text[:threshold], " ")
		if splitAt == -1 {
			splitAt = threshold
		}
		result.WriteString(text[:splitAt] + "\n" + indent)
		text = text[splitAt+1:]
	}
	result.WriteString(text)
	return result.String()
}
