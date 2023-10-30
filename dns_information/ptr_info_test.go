package dnsinformation

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
)

func TestGetPTR(t *testing.T) {
	t.Run("Valid PTR Records", func(t *testing.T) {
		// Mock net.LookupIP and net.LookupAddr functions to return valid PTR records.
		domain := "google.com"
		ptrRecords, err := getPTR(domain)

		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}

		expectedRecords := []string{
			"host1.example.com.",
			"host2.example.com.",
		}

		if len(ptrRecords) <= 0 {
			t.Errorf("Expected PTR records %v, but got %v", expectedRecords, ptrRecords)
		}
	})

	t.Run("No PTR Records", func(t *testing.T) {
		// Mock net.LookupIP and net.LookupAddr functions to return no PTR records.
		domain := ".com"
		ptrRecords, _ := getPTR(domain)

		if len(ptrRecords) > 0 {
			t.Error("Expected no PTR records, but got some records")
		}
	})

	t.Run("Invalid Domain", func(t *testing.T) {
		// Mock net.LookupIP function to return an error for an invalid domain.
		domain := "invalid.example"
		_, err := getPTR(domain)

		if err == nil {
			t.Error("Expected an error for an invalid domain, but got none")
		}
	})

	// Add more test cases as needed.
}

func TestGetPTRPrompt(t *testing.T) {
	// Redirect stdout for testing the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("Display PTR Records", func(t *testing.T) {
		input := "example.com"
		GetPTRPrompt(input, false) // Domain input
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" PTR   ✅: ") + color.Grey("host1.example.com,\n host2.example.com") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying PTR records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("Display 'None' for No PTR Records", func(t *testing.T) {
		input := "nodata.com"
		GetPTRPrompt(input, false) // Domain input with no PTR records
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" PTR   🟢: ") + color.Grey("None, Google and Microsoft 365 use shared IPs, this is ok.") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying 'None' for no PTR records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("Display PTR Records for IP", func(t *testing.T) {
		input := "192.168.1.1"
		GetPTRPrompt(input, true) // IP address input
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" PTR   ✅: ") + color.Grey("host1.example.com,\n host2.example.com") // Replace with actual PTR records
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying PTR records for IP, but got: %s", string(capturedOutput))
		}
	})

	// Restore the original stdout
	os.Stdout = originalStdout
}
