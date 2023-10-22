package dnsinformation

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
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

func TestGetRegistrarPrompt(t *testing.T) {
	// Redirect stdout for testing the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("Display Registrar Information", func(t *testing.T) {
		input := "example.com"
		GetRegistrarPromt(input, false) // Domain input
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" Registrar: ") + color.Green("Example Registrar") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying registrar information, but got: %s", string(capturedOutput))
		}
	})

	t.Run("Display 'Unknown or Classified'", func(t *testing.T) {
		input := "invalid.example"
		GetRegistrarPromt(input, false) // Invalid domain input
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" Registrar: ") + color.Green("Unknown or ") + color.Yellow("Classified") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying 'Unknown or Classified', but got: %s", string(capturedOutput))
		}
	})

	t.Run("Display Registrar Information for IP", func(t *testing.T) {
		input := "192.168.1.1"
		GetRegistrarPromt(input, true) // IP address input
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" Registrar: ") + color.Green("Registrar for IP addresses not available") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output displaying registrar information for IP, but got: %s", string(capturedOutput))
		}
	})

	// Restore the original stdout
	os.Stdout = originalStdout
}
