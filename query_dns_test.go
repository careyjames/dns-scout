package main

import "testing"

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
