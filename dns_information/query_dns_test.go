package dnsinformation

import (
	"net"
	"reflect"
	"testing"

	"github.com/miekg/dns"
)

func TestQueryDNS(t *testing.T) {
	// Mock the DNS query with a test server and response
	testServer := "8.8.8.8:53" // Example DNS server for testing
	testDomain := "google.com"
	testRecordType := dns.TypeA
	testResponse := &dns.Msg{}
	testResponse.Answer = []dns.RR{
		&dns.A{
			Hdr: dns.RR_Header{
				Name:   testDomain,
				Rrtype: dns.TypeA,
			},
			A: net.ParseIP("192.168.1.1"),
		},
		&dns.A{
			Hdr: dns.RR_Header{
				Name:   testDomain,
				Rrtype: dns.TypeA,
			},
			A: net.ParseIP("192.168.1.2"),
		},
	}

	dns.HandleFunc(testDomain, func(w dns.ResponseWriter, r *dns.Msg) {
		w.WriteMsg(testResponse)
	})
	defer dns.HandleRemove(testDomain)

	t.Run("Valid DNS Query", func(t *testing.T) {
		records, err := QueryDNS(testDomain, testRecordType, testServer)

		if err != nil {
			t.Errorf("Expected no error, but got an error: %v", err)
		}

		expectedRecords := []string{
			"192.168.1.1",
			"192.168.1.2",
		}

		if reflect.DeepEqual(records, expectedRecords) {
			t.Errorf("Expected DNS records %v, but got %v", expectedRecords, records)
		}
	})

	t.Run("Invalid DNS Server", func(t *testing.T) {
		// Provide an invalid DNS server address
		invalidServer := "invalid_server"

		_, err := QueryDNS(testDomain, testRecordType, invalidServer)

		if err == nil {
			t.Error("Expected an error for an invalid DNS server, but got none")
		}
	})

	// Add more test cases as needed.
}
