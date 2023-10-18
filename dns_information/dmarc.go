package dnsinformation

import (
	"fmt"
	"strings"
	"testing"

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
		if hasDMARCRecod(record) {
			return record, nil
		}
	}
	return "", nil
}

func hasDMARCRecod(record string) bool {
	return (len(record) > 8 && record[:8] == DMARCValid) || (len(record) > 6 && record[:6] == "vDMARC")
}

// GetDMARCPrompt fetches the DMARC record for a given domain.
func GetDMARCPrompt(input string) {
	dmarc, _ := getDMARC(input)
	if dmarc != "" {
		if isValidDMARC(dmarc) {
			formattedDMARC := formatLongText(dmarc, 80, " ")
			fmt.Printf(color.Blue(" DMARC Record: ") + color.Green(formattedDMARC) + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" DMARC Record: ") + color.Green(dmarc[8:]) + constants.Newline)
		}
	} else {
		fmt.Printf(color.Blue(" DMARC Record: ") + color.Red("None") + constants.Newline)
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

func TestFormatLongText(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description string
		text        string
		threshold   int
		indent      string
		expected    string
	}{
		{
			description: "Text is shorter than the threshold",
			text:        "Short text",
			threshold:   20,
			indent:      "  ",
			expected:    "Short text",
		},
		{
			description: "Text exactly matches the threshold",
			text:        "Exactly 20 characters",
			threshold:   20,
			indent:      "  ",
			expected:    "Exactly 20 characters",
		},
		{
			description: "Text longer than the threshold",
			text:        "This is a long text that needs to be formatted for better readability.",
			threshold:   20,
			indent:      "  ",
			expected:    "This is a long text\n  that needs to be\n  formatted for better\n  readability.",
		},
		{
			description: "Text with no spaces to split",
			text:        "ThisIsALongTextWithNoSpacesToSplit",
			threshold:   20,
			indent:      "  ",
			expected:    "ThisIsALongTextWithNoSpacesToSplit",
		},
		{
			description: "Empty text",
			text:        "",
			threshold:   20,
			indent:      "  ",
			expected:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := formatLongText(tc.text, tc.threshold, tc.indent)
			if result != tc.expected {
				t.Errorf("Expected result:\n%s\nGot:\n%s", tc.expected, result)
			}
		})
	}
}
