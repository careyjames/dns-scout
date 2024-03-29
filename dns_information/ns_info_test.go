package dnsinformation

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
)

func TestGetNS(t *testing.T) {
	// Celliwig: Disable for LaunchPad
//	t.Run("Valid NS Records from Google DNS", func(t *testing.T) {
//		// Mock Google DNS to return NS records.
//		domain := "example.com"
//		nsRecords, err := GetNS(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"ns1.example.com.",
//			"ns2.example.com.",
//		}
//
//		if len(nsRecords) <= 0 {
//			t.Errorf("Expected NS records %v, but got %v", expectedRecords, nsRecords)
//		}
//	})

	// Celliwig: Disable for LaunchPad
//	t.Run("Valid NS Records from Cloudflare DNS", func(t *testing.T) {
//		// Mock Cloudflare DNS to return NS records.
//		domain := "example.com"
//		nsRecords, err := GetNS(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"ns3.example.com.",
//			"ns4.example.com.",
//		}
//
//		if len(nsRecords) <= 0 {
//			t.Errorf("Expected NS records %v, but got %v", expectedRecords, nsRecords)
//		}
//	})

	t.Run("No NS Records", func(t *testing.T) {
		// Mock both DNS servers to return no NS records.
		domain := ".com"
		nsRecords, _ := GetNS(domain)

		if len(nsRecords) > 0 {
			t.Error("Expected no NS records, but got some records")
		}
	})

	t.Run("Both DNS Queries Failed", func(t *testing.T) {
		// Mock both DNS servers to return errors.
		domain := ".com"
		_, err := GetNS(domain)

		if err == nil {
			t.Error("Expected an error for both DNS queries failing, but got none")
		}
	})

	// Add more test cases as needed.
}

func TestGetNSPrompt(t *testing.T) {
	// Redirect stdout for testing the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("Display NS Records", func(t *testing.T) {
		input := "example.com"
		GetNSPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" Name Servers: ") + color.Grey("ns1.example.com, ns2.example.com") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying NS records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("No NS Records", func(t *testing.T) {
		input := "nodata.com"
		GetNSPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" Name Servers: ") + color.Red("None") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output containing 'None' for no NS records, but got: %s", string(capturedOutput))
		}
	})

	// Restore the original stdout
	os.Stdout = originalStdout
}
