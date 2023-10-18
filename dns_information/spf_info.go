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
			fmt.Printf(color.Blue(" SPF Records: ") + coloredSPFRecord + constants.Newline)
		} else {
			coloredSPFRecord := colorCodeSPFRecord(spfRecord) // "No SPF record" will be red
			fmt.Printf(color.Blue(" SPF Records: ") + coloredSPFRecord + constants.Newline)
		}
	}
}

func colorCodeSPFRecord(record string) string {
	colorCode := constants.GreenColorEncoding

	if strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
		colorCode = constants.RedColorEncoding
		if IsValidSPF(record) {
			colorCode = constants.GreenColorEncoding
		}
	}

	if !IsValidSPF(record) && strings.Contains(strings.ToLower(record), "spf") {
		colorCode = constants.RedColorEncoding
	}

	if record == " No SPF record" {
		colorCode = constants.RedColorEncoding
	}

	if strings.Contains(record, "-all") {
		record = strings.ReplaceAll(record, "-all", color.Yellow("-all"))
	} else if strings.Contains(record, "~all") {
		record = strings.ReplaceAll(record, "~all", color.Green("~all"))
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
