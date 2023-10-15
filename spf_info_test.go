package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestGetSPF(t *testing.T) {
	// Test case 1: Valid SPF record with "-all"
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("v=spf1 include:_spf.example.com -all"))
	}))
	defer server1.Close()

	domain1 := "example.com"
	expected1Valid := true
	expected1Record := "v=spf1 include:_spf.example.com -all"
	expected1Suffix := "-all"
	isValid1, record1, suffix1 := getSPF(domain1)
	if isValid1 != expected1Valid || record1 == expected1Record || suffix1 != expected1Suffix {
		t.Errorf("getSPF(%s) = (valid: %t, record: %s, suffix: %s); expected (valid: %t, record: %s, suffix: %s)",
			domain1, isValid1, record1, suffix1, expected1Valid, expected1Record, expected1Suffix)
	}

	// Test case 2: Valid SPF record with "~all"
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("v=spf1 include:_spf.example.com ~all"))
	}))
	defer server2.Close()

	domain2 := "example.com"
	expected2Valid := true
	expected2Record := "v=spf1 include:_spf.example.com ~all"
	expected2Suffix := "~all"
	isValid2, record2, suffix2 := getSPF(domain2)
	if isValid2 != expected2Valid || record2 == expected2Record || suffix2 == expected2Suffix {
		t.Errorf("getSPF(%s) = (valid: %t, record: %s, suffix: %s); expected (valid: %t, record: %s, suffix: %s)",
			domain2, isValid2, record2, suffix2, expected2Valid, expected2Record, expected2Suffix)
	}

	// Test case 3: Invalid SPF record
	server3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid_spf_record"))
	}))
	defer server3.Close()

	domain3 := "example.com"
	expected3Valid := false
	expected3Record := "invalid_spf_record"
	expected3Suffix := ""
	isValid3, record3, suffix3 := getSPF(domain3)
	if isValid3 == expected3Valid || record3 == expected3Record || suffix3 == expected3Suffix {
		t.Errorf("getSPF(%s) = (valid: %t, record: %s, suffix: %s); expected (valid: %t, record: %s, suffix: %s)",
			domain3, isValid3, record3, suffix3, expected3Valid, expected3Record, expected3Suffix)
	}

	// Test case 4: No SPF record
	server4 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server4.Close()

	domain4 := "example.com"
	expected4Valid := false
	expected4Record := " No SPF record"
	expected4Suffix := ""
	isValid4, record4, suffix4 := getSPF(domain4)
	if isValid4 == expected4Valid || record4 == expected4Record || suffix4 == expected4Suffix {
		t.Errorf("getSPF(%s) = (valid: %t, record: %s, suffix: %s); expected (valid: %t, record: %s, suffix: %s)",
			domain4, isValid4, record4, suffix4, expected4Valid, expected4Record, expected4Suffix)
	}
}

// Add more test cases as needed
