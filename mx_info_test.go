package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGetMXPrompt(t *testing.T) {
	// Test case 1: valid input with MX records
	input1 := "example.com"
	expected1 := []string{"mx1.example.com", "mx2.example.com"}
	mx1, _ := getMX(input1)
	if reflect.DeepEqual(mx1, expected1) {
		t.Errorf("getMXPrompt(%q) = %q; expected %q", input1, mx1, expected1)
	}
	var buf bytes.Buffer
	GetMXPrompt(input1)
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
	GetMXPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m Mail Servers: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
		t.Errorf("getMXPrompt(%q) output = %q; expected %q", input2, output2, expectedOutput2)
	}
}

func TestGetMX(t *testing.T) {
	// Test case 1: Valid domain with MX records
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("10 mx.example.com.\n20 mx2.example.com."))
	}))
	defer server1.Close()

	domain1 := "example.com"
	expected1 := []string{"10 mx.example.com.", "20 mx2.example.com."}
	result1, err1 := getMX(domain1)
	if err1 != nil {
		t.Errorf("getMX(%s) returned an error: %v", domain1, err1)
	}
	if compareStringSlices(result1, expected1) {
		t.Errorf("getMX(%s) = %v; expected %v", domain1, result1, expected1)
	}

	// Test case 2: Valid domain with no MX records
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server2.Close()

	domain2 := "example2.com"
	expected2 := []string{}
	result2, err2 := getMX(domain2)
	if err2 != nil {
		t.Errorf("getMX(%s) returned an error: %v", domain2, err2)
	}
	if compareStringSlices(result2, expected2) {
		t.Errorf("getMX(%s) = %v; expected %v", domain2, result2, expected2)
	}

	// Add more test cases as needed
}
