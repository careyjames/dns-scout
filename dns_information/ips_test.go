package dnsinformation

import (
	"bytes"
	"net"
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

func TestIPsToStrings(t *testing.T) {
	// Define test cases
	testCases := []struct {
		description        string
		ips                []net.IP
		expectedStringList []string
	}{
		{
			description: "Convert IPv4 addresses to strings",
			ips: []net.IP{
				net.IPv4(192, 0, 2, 1),
				net.IPv4(203, 0, 113, 2),
			},
			expectedStringList: []string{"192.0.2.1", "203.0.113.2"},
		},
		{
			description: "Convert IPv6 addresses to strings",
			ips: []net.IP{
				net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
				net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7335"),
			},
			expectedStringList: []string{
				"2001:db8:85a3::8a2e:370:7334",
				"2001:db8:85a3::8a2e:370:7335",
			},
		},
		{
			description:        "Empty input",
			ips:                []net.IP{},
			expectedStringList: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := ipsToStrings(tc.ips)
			if !stringSlicesEqual(result, tc.expectedStringList) {
				t.Errorf("Expected result: %v, got: %v", tc.expectedStringList, result)
			}
		})
	}
}

// Helper function to compare two string slices
func stringSlicesEqual(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
