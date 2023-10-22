package dnsinformation

import (
	"testing"
)

func TestGetTXTRecordNSLookup(t *testing.T) {
	t.Run("Valid TXT Records", func(t *testing.T) {
		// Mock the net.LookupTXT function to return TXT records.
		domain := "google.com"
		txtRecords, err := GetTXTRecordNSLookup(domain)

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
		// Mock the net.LookupTXT function to return an empty result.
		domain := "nodata.com"
		txtRecords, _ := GetTXTRecordNSLookup(domain)

		if len(txtRecords) > 0 {
			t.Error("Expected no TXT records, but got some records")
		}
	})

	t.Run("Non-Existent Domain", func(t *testing.T) {
		// Mock the net.LookupTXT function to return an error for a non-existent domain.
		domain := ".com"
		_, err := GetTXTRecordNSLookup(domain)

		if err == nil {
			t.Error("Expected an error for a non-existent domain, but got none")
		}
	})

	// Add more test cases as needed.
}

func TestGetDMARCRecordNSLookup(t *testing.T) {
	t.Run("Valid DMARC Records", func(t *testing.T) {
		// Mock the net.LookupTXT function to return DMARC records.
		domain := "google.com"
		dmarcRecords, err := GetDMARCRecordNSLookup(domain)

		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}

		expectedRecords := []string{
			"DMARC Record 1",
			"DMARC Record 2",
		}

		if len(dmarcRecords) <= 0 {
			t.Errorf("Expected DMARC records %v, but got %v", expectedRecords, dmarcRecords)
		}
	})

	t.Run("No DMARC Records", func(t *testing.T) {
		// Mock the net.LookupTXT function to return an empty result.
		domain := ".com"
		dmarcRecords, _ := GetDMARCRecordNSLookup(domain)

		if len(dmarcRecords) > 0 {
			t.Error("Expected no DMARC records, but got some records")
		}
	})

	t.Run("Non-Existent Domain", func(t *testing.T) {
		// Mock the net.LookupTXT function to return an error for a non-existent domain.
		domain := ".com"
		_, err := GetDMARCRecordNSLookup(domain)

		if err == nil {
			t.Error("Expected an error for a non-existent domain, but got none")
		}
	})

	// Add more test cases as needed.
}
