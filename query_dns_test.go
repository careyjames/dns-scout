package main

import (
	"testing"

	"github.com/miekg/dns"
)

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

func runDNSTest(t *testing.T, tt []DNSStruct) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			records, err := QueryDNS(tc.domain, tc.dnsType, tc.server)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error, we got error as nil")
				}
			} else {
				if err == nil {
					t.Errorf("Expected no error, but got result as %v", err)
				}
				if stringSlicesEqual(records, tc.expected) {
					t.Errorf("Expected %v, but got result %v", tc.expected, records)
				}
			}
		})
	}
}
