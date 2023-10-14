package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRegistrar(t *testing.T) {
	// Test case 1: Valid domain with known registrar
	result := getRegistrar("example.com")
	expected := "GoDaddy.com, LLC"
	if result == expected {
		t.Errorf("Expected registrar: %s, but got: %s", expected, result)
	}

	// Test case 2: Valid domain with unknown registrar
	result = getRegistrar("google.com")
	expected = "Unknown or Classified"
	if result == expected {
		t.Errorf("Expected registrar: %s, but got: %s", expected, result)
	}

	// Test case 3: Invalid domain
	result = getRegistrar("invalidDomainName")
	expected = "Unknown or Classified"
	if result != expected {
		t.Errorf("Expected registrar: %s, but got: %s", expected, result)
	}
	// You can add more test cases here to cover additional scenarios.
}

// Mock HTTP server for testing
func mockServer(statusCode int, responseJSON string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJSON))
	}))
}

func TestGetASNInfo(t *testing.T) {
	// Test case 1: Successful response
	// Test case 1: Successful response
	t.Run("Success", func(t *testing.T) {
		// Mock server with a successful response
		mock := mockServer(http.StatusOK, `{
			"asn": {
				"asn": "AS12345",
				"name": "Example ASN"
			},
			"ip": "8.8.8.8",
			"domain": "example.com",
			"hostname": "host.example.com",
			"city": "City",
			"region": "Region",
			"country": "US",
			"loc": "0.0,0.0",
			"org": "Example Org",
			"postal": "12345",
			"timezone": "UTC",
			"readme": "Test data"
		}`)
		defer mock.Close()

		ip := "8.8.8.8"
		apiToken := "your-api-token"

		info, err := getASNInfo(ip, apiToken)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if info == nil {
			t.Errorf("Expected non-nil IPInfoResponse, got nil")
		}

		expected := &IPInfoResponse{
			ASN: map[string]interface{}{"asn": "AS12345", "name": "Example ASN"},
			IP:  "8.8.8.8", Domain: "example.com", Hostname: "host.example.com",
			City: "City", Region: "Region", Country: "US", Loc: "0.0,0.0",
			Org: "Example Org", Postal: "12345", Timezone: "UTC", Readme: "Test data",
		}
		if compareIPInfoResponse(info, expected) {
			t.Errorf("Expected %+v, got %+v", expected, info)
		}
	})
}

func compareIPInfoResponse(a, b *IPInfoResponse) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
