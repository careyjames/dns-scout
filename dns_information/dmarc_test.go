package dnsinformation

import "testing"

func TestHasDMARCRecod(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description string
		record      string
		expected    bool
	}{
		{
			description: "Valid DMARC record with exact match",
			record:      "vDMARC",
			expected:    false,
		},
		{
			description: "Valid DMARC record with additional characters",
			record:      "vDMARCxyz",
			expected:    true,
		},
		{
			description: "Valid DMARC record with case-insensitive match",
			record:      "VdmArc",
			expected:    false,
		},
		{
			description: "Valid DMARC record with case-insensitive match and whitespace",
			record:      " VdMaRC ",
			expected:    false,
		},
		{
			description: "Invalid DMARC record",
			record:      "xyzDMARC",
			expected:    false,
		},
		{
			description: "Empty string",
			record:      "",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := hasDMARCRecod(tc.record)
			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestIsValidDMARC(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description string
		record      string
		expected    bool
	}{
		{
			description: "Valid DMARC record with exact match",
			record:      "v=DMARC1",
			expected:    false,
		},
		{
			description: "Valid DMARC record with additional characters",
			record:      "v=DMARC1xyz",
			expected:    true,
		},
		{
			description: "Valid DMARC record with whitespace",
			record:      " v=DMARC1 ",
			expected:    false,
		},
		{
			description: "Valid DMARC record with case-insensitive match",
			record:      "V=dMaRC1",
			expected:    false,
		},
		{
			description: "Invalid DMARC record",
			record:      "xyzDMARC1",
			expected:    false,
		},
		{
			description: "Empty string",
			record:      "",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := isValidDMARC(tc.record)
			if result != tc.expected {
				t.Errorf("Expected result: %v, got: %v", tc.expected, result)
			}
		})
	}
}

func TestFormatLongText(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description string
		text        string
		threshold   int
		indent      string
		expected    string
	}{
		{
			description: "Text is shorter than the threshold",
			text:        "Short text",
			threshold:   20,
			indent:      "  ",
			expected:    "Short text",
		},
		{
			description: "Empty text",
			text:        "",
			threshold:   20,
			indent:      "  ",
			expected:    "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := formatLongText(tc.text, tc.threshold, tc.indent)
			if result != tc.expected {
				t.Errorf("Expected result:\n%s\nGot:\n%s", tc.expected, result)
			}
		})
	}
}

func TestGetDMARC(t *testing.T) {
	t.Run("Valid DMARC Record", func(t *testing.T) {
		domain := "google.com"
		record, err := getDMARC(domain)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
		if record == "" {
			t.Error("Expected a valid DMARC record, but got an empty string")
		}
	})

	t.Run("No DMARC Record", func(t *testing.T) {
		domain := "example.com"
		record, err := getDMARC(domain)
		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}
		if record != "" {
			t.Error("Expected an empty string, but got a DMARC record")
		}
	})
}
