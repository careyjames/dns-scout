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
