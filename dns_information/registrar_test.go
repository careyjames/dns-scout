package dnsinformation

import (
	"testing"
)

func TestGetRegistrar(t *testing.T) {
	t.Run("Valid Registrar Information", func(t *testing.T) {
		// Mock the whois.Whois function to return valid registrar information.
		domain := "example.com"
		registrar := getRegistrar(domain)

		expectedRegistrar := "Example Registrar"
		if len(registrar) <= 0 {
			t.Errorf("Expected registrar '%s', but got '%s'", expectedRegistrar, registrar)
		}
	})

	t.Run("Invalid Domain", func(t *testing.T) {
		// Mock the whois.Whois function to return an error for an invalid domain.
		domain := "invalid.example"
		registrar := getRegistrar(domain)

		if registrar != ErrorMessage {
			t.Errorf("Expected an error message, but got '%s'", registrar)
		}
	})

	t.Run("Invalid WHOIS Data", func(t *testing.T) {
		// Mock the whois.Whois function to return invalid WHOIS data.
		domain := "error.example"
		registrar := getRegistrar(domain)

		if registrar != ErrorMessage {
			t.Errorf("Expected an error message, but got '%s'", registrar)
		}
	})

	t.Run("Registrar Not Found", func(t *testing.T) {
		// Mock the whois.Whois function to return WHOIS data without registrar information.
		domain := "noinfo.example"
		registrar := getRegistrar(domain)

		if registrar != ErrorMessage {
			t.Errorf("Expected an error message, but got '%s'", registrar)
		}
	})

	// Add more test cases as needed.
}
