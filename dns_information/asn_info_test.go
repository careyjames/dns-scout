package dnsinformation

import "testing"

func TestGetURL(t *testing.T) {
	expectedURL := "https://ipinfo.io/"

	// Call the function and get the actual URL
	actualURL := getURL()

	// Compare the expected and actual values
	if actualURL != expectedURL {
		t.Errorf("Expected URL: %s, but got: %s", expectedURL, actualURL)
	}
}
