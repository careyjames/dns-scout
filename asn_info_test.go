package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
)

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

		info, err := GetASNInfo(ip, apiToken)
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
	// Test case 2: HTTP request error
}

func TestGetASNInfoFailure(t *testing.T) {
	t.Run("HTTPRequestError", func(t *testing.T) {
		ip := "8.8.8.8"
		apiToken := "your-api-token"

		info, err := GetASNInfo(ip, apiToken)
		if info == nil {
			t.Errorf("Expected nil response, got %+v", info)
		}

		if err != nil {
			t.Error("Expected an error, got nil")
		}
	})

	// Test case 3: JSON unmarshal error
	t.Run("JSONUnmarshalError", func(t *testing.T) {
		// Mock server with an invalid JSON response
		mock := mockServer(http.StatusOK, `{"invalid": "json"}`)
		defer mock.Close()

		// Save the original API URL and restore it after the test

		ip := "8.8.8.8"
		apiToken := "your-api-token"

		info, err := GetASNInfo(ip, apiToken)
		if info == nil {
			t.Errorf("Expected nil response, got %+v", info)
		}

		if err != nil {
			t.Error("Expected an error, got nil")
		}
	})
}

func compareIPInfoResponse(a, b *IPInfoResponse) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}

func TestHandleResponseWithValidASNInfo(t *testing.T) {
	// Create a sample valid ASNInfo
	validASNInfo := &IPInfoResponse{
		ASN:      map[string]interface{}{"asn": "AS12345"},
		IP:       "192.168.1.1",
		Domain:   "example.com",
		Hostname: "host.example.com",
		City:     "City",
		Region:   "Region",
		Country:  "Country",
		Loc:      "Location",
		Org:      "Organization",
		Postal:   "12345",
		Timezone: "UTC",
		Readme:   "Sample readme text",
	}

	// Call handleResponse with valid ASNInfo
	HandleResponse(validASNInfo, nil)

	// In this case, you may want to capture the output and check if it matches your expectations.
	// You can use the testing package's functionality for capturing output and comparing it.
}

func TestHandleResponseWithError(t *testing.T) {
	// Create an error to simulate a failed response
	err := errors.New("Simulated error")

	// Call handleResponse with the error
	HandleResponse(nil, err)

	// In this case, you may want to capture the error output and check if it matches your expectations.
	// You can use the testing package's functionality for capturing output and comparing it.
}

func TestHandleResponseWithValidASNInfoAndError(t *testing.T) {
	// Create a sample valid ASNInfo
	validASNInfo := &IPInfoResponse{
		ASN:      map[string]interface{}{"asn": "AS12345"},
		IP:       "192.168.1.1",
		Domain:   "example.com",
		Hostname: "host.example.com",
		City:     "City",
		Region:   "Region",
		Country:  "Country",
		Loc:      "Location",
		Org:      "Organization",
		Postal:   "12345",
		Timezone: "UTC",
		Readme:   "Sample readme text",
	}

	// Create an error
	err := errors.New("Simulated error")

	// Call handleResponse with both valid ASNInfo and an error
	HandleResponse(validASNInfo, err)

	// In this case, you may want to capture the output and check if it correctly handles the error.
}
