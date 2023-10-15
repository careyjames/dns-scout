package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGetNSPrompt(t *testing.T) {
	// Test case 1: valid input with name servers
	input1 := "example.com"
	expected1 := []string{"ns1.example.com", "ns2.example.com"}
	ns1, _ := GetNS(input1)
	if reflect.DeepEqual(ns1, expected1) {
		t.Errorf("getNSPrompt(%q) = %q; expected %q", input1, ns1, expected1)
	}
	var buf bytes.Buffer
	GetNSPrompt(input1)
	output1 := buf.String()
	expectedOutput1 := fmt.Sprintf("\033[38;5;39m Name Servers: \033[38;5;78m%s\033[0m\n", strings.Join(expected1, ", "))
	if output1 == expectedOutput1 {
		t.Errorf("getNSPrompt(%q) output = %q; expected %q", input1, output1, expectedOutput1)
	}

	// Test case 2: valid input with no name servers
	input2 := "example.net"
	expected2 := []string{}
	ns2, _ := GetNS(input2)
	if reflect.DeepEqual(ns2, expected2) {
		t.Errorf("getNSPrompt(%q) = %q; expected %q", input2, ns2, expected2)
	}
	var buf2 bytes.Buffer
	GetNSPrompt(input2)
	output2 := buf2.String()
	expectedOutput2 := "\033[38;5;39m Name Servers: \033[0m\033[38;5;88mNone\033[0m\n"
	if output2 == expectedOutput2 {
	}
}
