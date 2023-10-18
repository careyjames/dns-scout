package dnsinformation

import (
	"fmt"
	"strings"
)

// getDMARC fetches the DMARC record for a given domain.
func getDMARC(domain string) (string, error) {
	txtRecords, err := GetTXT("_dmarc." + domain)
	if len(txtRecords) <= 0 {
		txtRecords, _ = GetDMARCRecordNSLookup(domain)
	}
	if err != nil {
		return "", err
	}
	for _, record := range txtRecords {
		if len(record) > 8 && record[:8] == "v=DMARC1" {
			return record, nil
		} else if record[:6] == "vDMARC" {
			return record, nil
		}
	}
	return "", nil
}

// GetDMARCPrompt fetches the DMARC record for a given domain.
func GetDMARCPrompt(input string) {
	dmarc, _ := getDMARC(input)
	if dmarc != "" {
		if isValidDMARC(dmarc) {
			formattedDMARC := formatLongText(dmarc, 80, " ")
			fmt.Printf("\033[38;5;39m DMARC Record:\033[0m\n")
			fmt.Printf("\033[38;5;78m %s\033[0m\n", formattedDMARC)
		} else {
			fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88m%s\033[0m\n", dmarc[8:])
		}
	} else {
		fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88mNone\033[0m\n")
	}
}

func isValidDMARC(record string) bool {
	if len(record) > 8 && record[:8] == "v=DMARC1" {
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
