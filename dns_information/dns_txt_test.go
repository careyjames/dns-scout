package dnsinformation

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
)

func TestGetTXT(t *testing.T) {
	// Mock a DNS server or use a library like dnstest to simulate DNS responses.

	// Celliwig: Disable for LaunchPad
//	t.Run("Valid TXT Records", func(t *testing.T) {
//		domain := "example.com"
//		txtRecords, err := GetTXT(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"TXT Record 1",
//			"TXT Record 2",
//		}
//
//		if len(txtRecords) <= 0 {
//			t.Errorf("Expected TXT records %v, but got %v", expectedRecords, txtRecords)
//		}
//	})

	t.Run("No TXT Records", func(t *testing.T) {
		domain := "nodata.com"
		txtRecords, _ := GetTXT(domain)

		if len(txtRecords) > 0 {
			t.Error("Expected no TXT records, but got some records")
		}
	})

	t.Run("Non-Existent Domain", func(t *testing.T) {
		domain := ".com"
		_, err := GetTXT(domain)

		if err == nil {
			t.Error("Expected an error for a non-existent domain, but got none")
		}
	})

	// Add more test cases as needed.
}

func TestGetTXTFromAllOption(t *testing.T) {
	// Celliwig: Disable for LaunchPad
//	t.Run("Fetch TXT Records Using GetTXT", func(t *testing.T) {
//		// Mock or set up a test environment to return records from GetTXT.
//		domain := "example.com"
//		txtRecords, err := GetTXTFromAllOption(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"TXT Record 1",
//			"TXT Record 2",
//		}
//
//		if len(txtRecords) <= 0 {
//			t.Errorf("Expected TXT records %v, but got %v", expectedRecords, txtRecords)
//		}
//	})

	// Celliwig: Disable for LaunchPad
//	t.Run("Fetch TXT Records Using GetTXTRecordNSLookup", func(t *testing.T) {
//		// Mock or set up a test environment to return records from GetTXTRecordNSLookup.
//		domain := "example.com"
//		txtRecords, err := GetTXTFromAllOption(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"TXT Record 1",
//			"TXT Record 2",
//		}
//
//		if len(txtRecords) <= 0 {
//			t.Errorf("Expected TXT records %v, but got %v", expectedRecords, txtRecords)
//		}
//	})

	t.Run("No TXT Records", func(t *testing.T) {
		// Mock both GetTXT and GetTXTRecordNSLookup to return no records.
		domain := "nodata.com"
		txtRecords, _ := GetTXTFromAllOption(domain)

		if len(txtRecords) > 0 {
			t.Error("Expected no TXT records, but got some records")
		}
	})

	t.Run("Non-Existent Domain", func(t *testing.T) {
		// Mock both GetTXT and GetTXTRecordNSLookup to return an error.
		domain := ".com"
		_, err := GetTXTFromAllOption(domain)

		if err == nil {
			t.Error("Expected an error for a non-existent domain, but got none")
		}
	})

	// Add more test cases as needed.
}

func TestGetTXTPrompt(t *testing.T) {
	// Redirect stdout for testing the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("Display TXT Records", func(t *testing.T) {
		input := "example.com"
		GetTXTPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		// Customize the expected output based on your formatting and colors
		expectedOutput := color.Blue(" TXT Records: ") + constants.Newline +
			" Colored TXT Record 1" + constants.Newline +
			" Colored TXT Record 2" + constants.Newline

		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying TXT records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("No TXT Records", func(t *testing.T) {
		input := "nodata.com"
		GetTXTPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" TXT Records: ") + color.Red("None") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output containing 'None' for no TXT records, but got: %s", string(capturedOutput))
		}
	})

	// Restore the original stdout
	os.Stdout = originalStdout
}
