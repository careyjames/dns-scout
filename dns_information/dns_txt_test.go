package dnsinformation

import (
	"testing"
)

func TestGetTXT(t *testing.T) {
	// Mock a DNS server or use a library like dnstest to simulate DNS responses.

	t.Run("Valid TXT Records", func(t *testing.T) {
		domain := "example.com"
		txtRecords, err := GetTXT(domain)

		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}

		expectedRecords := []string{
			"TXT Record 1",
			"TXT Record 2",
		}

		if len(txtRecords) <= 0 {
			t.Errorf("Expected TXT records %v, but got %v", expectedRecords, txtRecords)
		}
	})

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
