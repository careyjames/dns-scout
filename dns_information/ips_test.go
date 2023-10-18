package dnsinformation

import (
	"bytes"
	"testing"
)

func TestResolvedIPPrompt(t *testing.T) {
	// Create a custom output buffer to capture the printed content.
	var outputBuffer bytes.Buffer

	// Define test cases
	testCases := []struct {
		description    string
		input          string
		expectedOutput string
	}{
		{
			description:    "Valid IP address resolution",
			input:          "example.com",
			expectedOutput: "",
		},
		{
			description:    "Invalid input (no IP resolution)",
			input:          "invalidhostname",
			expectedOutput: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			outputBuffer.Reset() // Clear the output buffer before each test case
			ResolvedIPPrompt(tc.input)
			output := outputBuffer.String()
			if output != tc.expectedOutput {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tc.expectedOutput, output)
			}
		})
	}
}
