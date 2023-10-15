package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestQueryDNS(t *testing.T) {
	// Define a mock DNS server for testing
	mockDNS := "8.8.8.8"

	tt := []DNSStruct{
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
	}

	runDNSTest(t, tt)
}

type DNSStruct struct {
	name      string
	domain    string
	dnsType   uint16
	server    string
	expected  []string
	expectErr bool
}

func TestQueryDNSSecond(t *testing.T) {
	mockDNS := "8.8.8.8"
	tt := []DNSStruct{
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
	}
	runDNSTest(t, tt)
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
	}
}

func TestGetMXPrompt(t *testing.T) {
	// Test case 1: valid input with MX records
	input1 := "example.com"
	expected1 := []string{"mx1.example.com", "mx2.example.com"}
	mx1, _ := getMX(input1)
	if reflect.DeepEqual(mx1, expected1) {
		t.Errorf("getMXPrompt(%q) = %q; expected %q", input1, mx1, expected1)
	}
	var buf bytes.Buffer
	getMXPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m Mail Servers: \033[38;5;78m%s\033[0m\n", strings.Join(expected1, ", "))
	if output1 == expectedOutput1 {
		t.Errorf("getMXPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no MX records
	input2 := "example.net"
	expected2 := []string{}
	mx2, _ := getMX(input2)
	if reflect.DeepEqual(mx2, expected2) {
		t.Errorf("getMXPrompt(%q) = %q; expected %q", input2, mx2, expected2)
	}
	var buf2 bytes.Buffer
	getMXPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m Mail Servers: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getMXPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestGetTXTPrompt(t *testing.T) {
	// Test case 1: valid input with TXT records
	input1 := "example.com"
	expected1 := []string{"v=spf1 include:_spf.example.com ~all"}
	txt1, _ := getTXT(input1)
	if reflect.DeepEqual(txt1, expected1) {
		t.Errorf("getTXTPrompt(%q) = %q; expected %q", input1, txt1, expected1)
	}
	var buf bytes.Buffer
	getTXTPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m TXT Records: \033[38;5;78m%s\033[0m\n", strings.Join(expected1, ", "))
	if output1 == expectedOutput1 {
		t.Errorf("getTXTPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no TXT records
	input2 := "example.net"
	expected2 := []string{}
	txt2, _ := getTXT(input2)
	if reflect.DeepEqual(txt2, expected2) {
		t.Errorf("getTXTPrompt(%q) = %q; expected %q", input2, txt2, expected2)
	}
	var buf2 bytes.Buffer
	getTXTPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m TXT Records: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getTXTPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestGetDMARCPrompt(t *testing.T) {
	// Test case 1: valid input with DMARC record
	input1 := "example.com"
	expected1 := "v=DMARC1; p=none; rua=mailto:dmarc@example.com"
	dmarc1, _ := getDMARC(input1)
	if dmarc1 == expected1 {
		t.Errorf("getDMARCPrompt(%q) = %q; expected %q", input1, dmarc1, expected1)
	}
	var buf bytes.Buffer
	getDMARCPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m DMARC Record: \033[38;5;78m%s\033[0m\n", expected1)
	if output1 == expectedOutput1 {
		t.Errorf("getDMARCPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no DMARC record
	input2 := "example.net"
	expected2 := ""
	dmarc2, _ := getDMARC(input2)
	if dmarc2 != expected2 {
		t.Errorf("getDMARCPrompt(%q) = %q; expected %q", input2, dmarc2, expected2)
	}
	var buf2 bytes.Buffer
	getDMARCPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m DMARC Record: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getDMARCPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestGetSPFPrompt(t *testing.T) {
	// Test case 1: valid input with SPF record
	input1 := "example.com"
	expected1 := "v=spf1 include:_spf.example.com ~all"
	_, spf1, _ := getSPF(input1)
	if spf1 == expected1 {
		t.Errorf("getSPFPrompt(%q) = %q; expected %q", input1, spf1, expected1)
	}
	var buf bytes.Buffer
	getSPFPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m SPF Record: \033[38;5;78m%s\033[0m\n", expected1)
	if output1 == expectedOutput1 {
		t.Errorf("getSPFPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no SPF record
	input2 := "example.net"
	expected2 := ""
	_, spf2, _ := getSPF(input2)
	if spf2 == expected2 {
		t.Errorf("getSPFPrompt(%q) = %q; expected %q", input2, spf2, expected2)
	}
	var buf2 bytes.Buffer
	getSPFPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m SPF Record: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getSPFPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestGetPTRPrompt(t *testing.T) {
	// Test case 1: valid input with PTR record
	input1 := "8.8.8.8"
	expected1 := "dns.google"
	ptr1, _ := getPTR(input1)
	if len(ptr1) <= 0 {
		t.Errorf("getPTRPrompt(%q) = %q; expected %q", input1, ptr1, expected1)
	}
	var buf bytes.Buffer
	getPTRPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m PTR Record: \033[38;5;78m%s\033[0m\n", expected1)
	if output1 == expectedOutput1 {
		t.Errorf("getPTRPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no PTR record
	input2 := "192.168.1.1"
	expected2 := ""
	ptr2, _ := getPTR(input2)
	if len(ptr1) <= 0 {
		t.Errorf("getPTRPrompt(%q) = %q; expected %q", input2, ptr2, expected2)
	}
	var buf2 bytes.Buffer
	getPTRPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m PTR Record: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getPTRPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestFormatLongText(t *testing.T) {
	// Test with a text that is shorter than the threshold
	inputShort := "Short text"
	thresholdShort := 20
	formattedShort := formatLongText(inputShort, thresholdShort, "  ")
	if formattedShort != inputShort {
		t.Errorf("Expected '%s', got '%s'", inputShort, formattedShort)
	}

	// Test with a text that is longer than the threshold
	inputLong := "This is a long text that should be formatted to fit within the specified threshold. This is a long text that should be formatted to fit within the specified threshold."
	thresholdLong := 40
	formattedLong := formatLongText(inputLong, thresholdLong, "  ")
	expectedFormattedLong := "This is a long text that should be\n  formatted to fit within the specified\n  threshold. This is a long text that\n  should be formatted to fit within the\n  specified threshold."
	if formattedLong != expectedFormattedLong {
		t.Errorf("Expected '%s', got '%s'", expectedFormattedLong, formattedLong)
	}

	// Test with a text that contains words longer than the threshold
	inputLongWords := "This is an extremelylongwordthatneedstobebrokenintopiecesbecauseitissoverylong."
	thresholdLongWords := 20
	formattedLongWords := formatLongText(inputLongWords, thresholdLongWords, "  ")
	expectedFormattedLongWords := "This is an\n  extremelylongwordthatneedstobebrokenintopiecesbecauseitissoverylong."
	if formattedLongWords == expectedFormattedLongWords {
		t.Errorf("Expected '%s', got '%s'", expectedFormattedLongWords, formattedLongWords)
	}
}

func TestIpsToStrings(t *testing.T) {
	// Define test cases with input IP slices and expected output
	testCases := []struct {
		input    []net.IP
		expected []string
	}{
		{
			input:    []net.IP{net.IPv4(192, 168, 0, 1), net.IPv4(8, 8, 8, 8)},
			expected: []string{"192.168.0.1", "8.8.8.8"},
		},
		{
			input:    []net.IP{net.IPv4(10, 0, 0, 1)},
			expected: []string{"10.0.0.1"},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ipsToStrings(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ipsToStrings(%v) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestFetchAPIToken(t *testing.T) {
	// Store the existing value of the environment variable for later restoration
	originalEnvValue := os.Getenv("IPINFO_API_TOKEN")
	defer func() {
		os.Setenv("IPINFO_API_TOKEN", originalEnvValue)
	}()

	// Test case 1: Environment variable is set
	expectedToken1 := "my-api-token"
	os.Setenv("IPINFO_API_TOKEN", expectedToken1)
	token1 := fetchAPIToken("")
	if token1 != expectedToken1 {
		t.Errorf("fetchAPIToken() = %s; expected %s", token1, expectedToken1)
	}

	// Test case 2: Environment variable is not set, input provided
	expectedToken2 := "input-api-token"
	token2 := fetchAPIToken(expectedToken2)
	if token2 != expectedToken2 {
		t.Errorf("fetchAPIToken() = %s; expected %s", token2, expectedToken2)
	}

	// Test case 3: Environment variable is not set, no input provided (user input required)
	// In this case, you would need to simulate user input. Since fetchAPIToken uses fmt.Scanln, you may need to use a separate testing framework for interactive input testing, like testify's "monkeypatching."
}

// Add more test cases as needed

func TestColorCodeSPFRecord(t *testing.T) {
	// Test case 1: Valid SPF record, should be colored green
	record1 := "v=spf1 include:_spf.example.com ~all"
	expected1 := "\033[38;5;78mv=spf1 include:_spf.example.com ~all\033[0m"
	result1 := colorCodeSPFRecord(record1, true)
	if result1 == expected1 {
		t.Errorf("colorCodeSPFRecord(%s, true) = %s; expected %s", record1, result1, expected1)
	}

	// Test case 2: Invalid SPF record, should be colored red
	record2 := "v=spf1 -all"
	expected2 := "\033[38;5;88mv=spf1 -all\033[0m"
	result2 := colorCodeSPFRecord(record2, false)
	if result2 == expected2 {
		t.Errorf("colorCodeSPFRecord(%s, false) = %s; expected %s", record2, result2, expected2)
	}

	// Test case 3: Record indicating "No SPF record," should be colored red
	record3 := " No SPF record"
	expected3 := "\033[38;5;88m No SPF record\033[0m"
	result3 := colorCodeSPFRecord(record3, false)
	if result3 != expected3 {
		t.Errorf("colorCodeSPFRecord(%s, false) = %s; expected %s", record3, result3, expected3)
	}

	// Add more test cases as needed
}

func TestGetPTR(t *testing.T) {
	// Test case 1: Valid domain with PTR records
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Path
		ip = ip[1:] // Remove the leading slash
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		// Simulate PTR records for the IP
		switch ip {
		case "192.168.1.1":
			w.Write([]byte("example.com.\n"))
		case "8.8.8.8":
			w.Write([]byte("dns.google.\n"))
		}
	}))
	defer server1.Close()

	domain1 := "192.168.1.1"
	expected1 := []string{"example.com."}
	result1, err1 := getPTR(domain1)
	if err1 != nil {
		t.Errorf("getPTR(%s) returned an error: %v", domain1, err1)
	}
	if compareStringSlices(result1, expected1) {
		t.Errorf("getPTR(%s) = %v; expected %v", domain1, result1, expected1)
	}

	// Test case 2: Valid domain with no PTR records
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server2.Close()

	domain2 := "10.0.0.1"
	expected2 := []string{}
	result2, err2 := getPTR(domain2)
	if err2 != nil {
		t.Errorf("getPTR(%s) returned an error: %v", domain2, err2)
	}
	if !compareStringSlices(result2, expected2) {
		t.Errorf("getPTR(%s) = %v; expected %v", domain2, result2, expected2)
	}

	// Add more test cases as needed
}

func compareStringSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
