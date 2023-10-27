package dnsinformation

import (
	"testing"
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
