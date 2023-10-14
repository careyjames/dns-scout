package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/miekg/dns"
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
	// Test case 2: HTTP request error
	t.Run("HTTPRequestError", func(t *testing.T) {
		ip := "8.8.8.8"
		apiToken := "your-api-token"

		info, err := getASNInfo(ip, apiToken)
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

		info, err := getASNInfo(ip, apiToken)
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

func TestQueryDNS(t *testing.T) {
	// Define a mock DNS server for testing
	mockDNS := "8.8.8.8"

	tt := []struct {
		name      string
		domain    string
		dnsType   uint16
		server    string
		expected  []string
		expectErr bool
	}{
		{
			name:      "Valid A record query",
			domain:    "example.com",
			dnsType:   dns.TypeA,
			server:    mockDNS,
			expected:  []string{"93.184.216.34"},
			expectErr: false,
		},
		{
			name:      "Valid NS record query",
			domain:    "example.com",
			dnsType:   dns.TypeNS,
			server:    mockDNS,
			expected:  []string{"a.iana-servers.net", "b.iana-servers.net"},
			expectErr: false,
		},
		{
			name:      "Valid MX record query",
			domain:    "example.com",
			dnsType:   dns.TypeMX,
			server:    mockDNS,
			expected:  []string{"0 aspmx.l.google.com", "5 alt1.aspmx.l.google.com"},
			expectErr: false,
		},
		{
			name:      "Invalid domain",
			domain:    "nonexistent.invalid",
			dnsType:   dns.TypeA,
			server:    mockDNS,
			expected:  nil,
			expectErr: true,
		},
		// Add more test cases as needed
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			records, err := queryDNS(tc.domain, tc.dnsType, tc.server)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err == nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if stringSlicesEqual(records, tc.expected) {
					t.Errorf("Expected %v, but got %v", tc.expected, records)
				}
			}
		})
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestGetNSPrompt(t *testing.T) {
	// Test case 1: valid input with name servers
	input1 := "example.com"
	expected1 := []string{"ns1.example.com", "ns2.example.com"}
	ns1, _ := getNS(input1)
	if reflect.DeepEqual(ns1, expected1) {
		t.Errorf("getNSPrompt(%q) = %q; expected %q", input1, ns1, expected1)
	}
	var buf bytes.Buffer
	getNSPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m Name Servers: \033[38;5;78m%s\033[0m\n", strings.Join(expected1, ", "))
	if output1 == expectedOutput1 {
		t.Errorf("getNSPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no name servers
	input2 := "example.net"
	expected2 := []string{}
	ns2, _ := getNS(input2)
	if reflect.DeepEqual(ns2, expected2) {
		t.Errorf("getNSPrompt(%q) = %q; expected %q", input2, ns2, expected2)
	}
	var buf2 bytes.Buffer
	getNSPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m Name Servers: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getNSPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}
