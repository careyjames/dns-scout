package dnsinformation

import (
	"fmt"
	"strings"
)

// getSPF fetches and analyzes the SPF record for a given domain.
func getSPF(domain string) (bool, string, string) {
	txtRecords, err := GetTXT(domain)
	if err != nil {
		return false, "Error fetching TXT records", ""
	}
	if len(txtRecords) <= 0 {
		txtRecords, _ = GetTXTRecordNSLookup(domain)
	}
	for _, record := range txtRecords {
		suffix := ""
		if strings.Contains(record, "-all") {
			suffix = "-all"
		} else if strings.Contains(record, "~all") {
			suffix = "~all"
		}

		if strings.HasPrefix(record, "v=spf1") {
			return true, record, suffix
		} else if strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
			return false, record, suffix
		}
	}
	return false, " No SPF record", ""
}

// GetSPFPrompt is prompt for spf
func GetSPFPrompt(input string) {
	spfValid, spfRecord, _ := getSPF(input)
	txt, _ := GetTXT(input)
	countTxt := 0
	if len(txt) > 1 {
		for _, record := range txt {
			if strings.Contains(strings.ToLower(record), "spf") {
				countTxt = countTxt + 1
			}
		}
	}
	if countTxt > 1 {
		fmt.Printf("\033[38;5;39m SPF Records: \033[0m\033[38;5;88mCan't have two SPF!\033[0m\n")
	} else {
		if spfValid || spfRecord != "No SPF record" {
			coloredSPFRecord := colorCodeSPFRecord(spfRecord, spfValid)
			fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
		} else {
			coloredSPFRecord := colorCodeSPFRecord(spfRecord, false) // "No SPF record" will be red
			fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
		}
	}
}

func colorCodeSPFRecord(record string, valid bool) string {
	colorCode := "\033[38;5;78m" // Default to green for non-SPF records

	if strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
		colorCode = "\033[38;5;88m" // Default to red for malformed or misspelled SPF
		if valid {
			colorCode = "\033[38;5;78m" // Green for valid SPF
		}
	}

	if !valid && strings.Contains(strings.ToLower(record), "spf") {
		colorCode = "\033[38;5;88m" // Green for invalid SPF
	}

	if record == " No SPF record" {
		colorCode = "\033[38;5;88m" // Red for "No SPF record"
	}

	if strings.Contains(record, "-all") {
		record = strings.ReplaceAll(record, "-all", "\033[38;5;222m-all\033[0m")
	} else if strings.Contains(record, "~all") {
		record = strings.ReplaceAll(record, "~all", "\033[38;5;78m~all\033[0m")
	}

	return fmt.Sprintf("%s%s\033[0m", colorCode, record)
}
