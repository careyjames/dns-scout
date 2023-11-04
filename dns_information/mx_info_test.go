package dnsinformation

import (
	"testing"
)

func TestGetMX(t *testing.T) {
	// Celliwig: Disable for LaunchPad
//	t.Run("Valid MX Records", func(t *testing.T) {
//		// Mock a DNS server or set up a test environment to return MX records.
//		domain := "example.com"
//		mxRecords, err := getMX(domain)
//
//		if err != nil {
//			t.Errorf("Expected no error, but got an error: %v", err)
//		}
//
//		expectedRecords := []string{
//			"mail.example.com.",
//			"backup.example.com.",
//		}
//
//		if len(mxRecords) <= 0 {
//			t.Errorf("Expected MX records %v, but got %v", expectedRecords, mxRecords)
//		}
//	})

	t.Run("No MX Records", func(t *testing.T) {
		// Mock the DNS server to return no MX records.
		domain := ".com"
		mxRecords, _ := getMX(domain)

		if len(mxRecords) > 0 {
			t.Error("Expected no MX records, but got some records")
		}
	})

	t.Run("Non-Existent Domain", func(t *testing.T) {
		// Mock the DNS server to return an error for a non-existent domain.
		domain := ".com"
		_, err := getMX(domain)

		if err == nil {
			t.Error("Expected an error for a non-existent domain, but got none")
		}
	})

	// Add more test cases as needed.
}
