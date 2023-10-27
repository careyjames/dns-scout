package dnsinformation

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/careyjames/DNS-Scout/dto"
)

func TestGetURL(t *testing.T) {
	expectedURL := "https://ipinfo.io/"

	// Call the function and get the actual URL
	actualURL := getURL()

	// Compare the expected and actual values
	if actualURL != expectedURL {
		t.Errorf("Expected URL: %s, but got: %s", expectedURL, actualURL)
	}
}

func TestIPInfoResponseUnmarshal(t *testing.T) {
	t.Run("Unmarshal Valid JSON", func(t *testing.T) {
		validJSON := `{
            "asn": {
                "asn": "AS12345",
                "name": "Example ASN"
            },
            "ip": "192.168.1.1",
            "domain": "example.com",
            "hostname": "host.example.com",
            "city": "City",
            "region": "Region",
            "country": "Country",
            "loc": "0.0000,0.0000",
            "org": "Example Organization",
            "postal": "12345",
            "timezone": "UTC",
            "readme": "This is a test"
        }`

		var response *dto.IPInfoResponse
		err := json.Unmarshal([]byte(validJSON), &response)

		if err != nil {
			t.Errorf("Expected no error when unmarshaling valid JSON, but got an error: %v", err)
		}

		expectedASN := map[string]interface{}{
			"asn":  "AS12345",
			"name": "Example ASN",
		}
		if !reflect.DeepEqual(response.ASN, expectedASN) {
			t.Errorf("Expected ASN %v, but got %v", expectedASN, response.ASN)
		}

		expectedIP := "192.168.1.1"
		if response.IP != expectedIP {
			t.Errorf("Expected IP %s, but got %s", expectedIP, response.IP)
		}
		// Add assertions for other fields in the response as needed
	})

	// Add more test cases as needed.
}
