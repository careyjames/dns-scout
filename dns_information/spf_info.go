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
		fmt.Printf(color.Blue(" SPF Record: ❌ ") + color.Red("There can only be ONE!") + constants.Newline)
	} else {
		if spfValid {
			coloredSPFRecord := "✅ " + spfRecord // Ensuring valid SPF records are green
			fmt.Printf(color.Blue(" SPF Record: ") + coloredSPFRecord + constants.Newline)
		} else {
			coloredSPFRecord := "❌ " + spfRecord // Ensuring invalid SPF records are red
			fmt.Printf(color.Blue(" SPF Record: ") + coloredSPFRecord + constants.Newline)
		}
	}
}

func colorCodeSPFRecord(record string) string {
	colorCode := constants.GreenColorEncoding // Default color is green

	// Check if it's a valid SPF record
	if IsValidSPF(record) {
		return fmt.Sprintf("%s%s\033[0m", colorCode, colorCodeSPFModifiers(record))
	}

	// Check if it's a TXT record (not SPF)
	if strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
		colorCode = constants.RedColorEncoding // Set color to red for invalid SPF records
		return fmt.Sprintf("%s%s\033[0m", colorCode, colorCodeSPFModifiers(record))
	}

	// For other TXT records, keep them green
	return fmt.Sprintf("%s%s\033[0m", constants.GreenColorEncoding, record)
}

func colorCodeSPFModifiers(record string) string {
	if strings.Contains(record, "-all") {
		record = strings.ReplaceAll(record, "-all", color.Yellow("-all"))
	} else if strings.Contains(record, "~all") {
		record = strings.ReplaceAll(record, "~all", color.Green("~all"))
	}
	return record
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

// IsValidSPF is valid spf
func IsValidSPF(record string) bool {
	return strings.HasPrefix(record, "v=spf1")
}

// HasInvalidSPFRecord is invalid spf
func HasInvalidSPFRecord(record string) bool {
	return strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all")
}

// HasSPFRecord has spf
func HasSPFRecord(record string) bool {
	return strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all")
}
