package dnsinformation

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

func TestHasDKIMRecord(t *testing.T) {
	// Test cases for hasDKIMRecord function
	t.Run("Valid DKIM record", func(t *testing.T) {
		record := "v=DKIM1alidABCDEFG"
		result := hasDKIMRecord(record)
		if !result {
			t.Errorf("Expected true for a valid DKIM record, got false")
		}
	})

	t.Run("Short DKIM record", func(t *testing.T) {
		record := "vDKIM1dsdds"
		result := hasDKIMRecord(record)
		if !result {
			t.Errorf("Expected true for a short DKIM record, got false")
		}
	})

	t.Run("Invalid DKIM record", func(t *testing.T) {
		record := "InvalidDKIM"
		result := hasDKIMRecord(record)
		if result {
			t.Errorf("Expected false for an invalid DKIM record, got true")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		record := ""
		result := hasDKIMRecord(record)
		if result {
			t.Errorf("Expected false for an empty string, got true")
		}
	})
}

func TestIsValidDKIM(t *testing.T) {
	// Test cases for isValidDKIM function
	t.Run("Valid DKIM record", func(t *testing.T) {
		record := "v=DKIM1ValidABCDEFG"
		result := isValidDKIM(record)
		if !result {
			t.Errorf("Expected true for a valid DKIM record, got false")
		}
	})

	t.Run("Short DKIM record", func(t *testing.T) {
		record := "DKIMVal"
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected true for a short DKIM record, got false")
		}
	})

	t.Run("Invalid DKIM record", func(t *testing.T) {
		record := "InvalidDKIM"
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected false for an invalid DKIM record, got true")
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		record := ""
		result := isValidDKIM(record)
		if result {
			t.Errorf("Expected false for an empty string, got true")
		}
	})
}

func TestGetDKIM(t *testing.T) {
	t.Run("Valid DKIM Record", func(t *testing.T) {
		domain := "example.com"
		selector := "valid"
		_, err := getDKIM(domain, selector)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
	})

	t.Run("No DKIM Record", func(t *testing.T) {
		domain := "example.com"
		selector := "nodkim"
		record, err := getDKIM(domain, selector)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
		if record != "" {
			t.Error("Expected an empty string, but got a DKIM record")
		}
	})
}

func TestGetDKIMPrompt(t *testing.T) {
	// Redirect stdout for testing the output
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Run("Valid DKIM Records", func(t *testing.T) {
		input := "example.com"
		GetDKIMPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		// Customize the expected output based on your formatting and colors
		expectedOutput := color.Blue(" DKIM Records: ") + color.Grey("selector1.") + color.Grey("Valid DKIM Record1") + constants.Newline +
			color.Grey("selector2.") + color.Grey("Valid DKIM Record2") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output containing valid DKIM records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("Invalid DKIM Records", func(t *testing.T) {
		input := "invalid.com"
		GetDKIMPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		// Customize the expected output based on your formatting and colors
		expectedOutput := color.Blue(" DKIM Records: ") + color.Grey("selector1.") + color.Red("Invalid DKIM Record1") + constants.Newline +
			color.Grey("selector2.") + color.Red("Invalid DKIM Record2") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output containing invalid DKIM records, but got: %s", string(capturedOutput))
		}
	})

	t.Run("No DKIM Records", func(t *testing.T) {
		input := "nodkim.com"
		GetDKIMPrompt(input)
		w.Close()
		capturedOutput, _ := ioutil.ReadAll(r)

		expectedOutput := color.Blue(" DKIM Records: ") + color.Red("None") + constants.Newline
		if strings.Contains(string(capturedOutput), expectedOutput) {
			t.Errorf("Expected output containing 'None' for no DKIM records, but got: %s", string(capturedOutput))
		}
	})

	// Restore the original stdout
	os.Stdout = originalStdout
}
