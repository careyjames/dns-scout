package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

// getSPF fetches and analyzes the SPF record for a given domain.
func getSPF(domain string) (bool, string) {
	txtRecords, err := GetTXTFromAllOption(domain)
	if err != nil {
		return false, "Error fetching TXT records"
	}

	for _, record := range txtRecords {
		if IsValidSPF(record) {
			return true, record
		} else if HasInvalidSPFRecord(record) {
			return false, record
		}
	}
	return false, " No SPF record"
}

// GetSPFPrompt is prompt for spf
func GetSPFPrompt(input string) {
	spfValid, spfRecord := getSPF(input)
	txt, _ := GetTXTFromAllOption(input)
	countTxt := totalSPFRecords(txt)

	if countTxt > 1 {
		fmt.Printf(color.Blue(" SPF Records: ") + color.Red("Can't have two SPF!") + constants.Newline)
	} else {
		if spfValid {
			coloredSPFRecord := colorCodeSPFRecord(spfRecord)
			fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
		} else {
			coloredSPFRecord := colorCodeSPFRecord(spfRecord) // "No SPF record" will be red
			fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
		}
	}
}

func colorCodeSPFRecord(record string) string {
	colorCode := "\033[38;5;78m" // Default to green for non-SPF records

	if strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
		colorCode = "\033[38;5;88m" // Default to red for malformed or misspelled SPF
		if IsValidSPF(record) {
			colorCode = "\033[38;5;78m" // Green for valid SPF
		}
	}

	if !IsValidSPF(record) && strings.Contains(strings.ToLower(record), "spf") {
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

func totalSPFRecords(records []string) int {
	countTxt := 0
	if len(records) > 1 {
		for _, record := range records {
			if strings.Contains(strings.ToLower(record), "spf") {
				countTxt = countTxt + 1
			}
		}
	}
	return countTxt
}

func IsValidSPF(record string) bool {
	return strings.HasPrefix(record, "v=spf1")
}

func HasInvalidSPFRecord(record string) bool {
	return strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all")
}

func HasSPFRecord(record string) bool {
	return strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all")
}
