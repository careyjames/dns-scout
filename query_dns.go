package main

import (
	"strings"

	"github.com/miekg/dns"
)

// QueryDNS performs a DNS query for a given domain and DNS record type.
func QueryDNS(domain string, dnsType uint16, server string) ([]string, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dnsType)

	r, _, err := c.Exchange(m, server)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, ans := range r.Answer {
		switch t := ans.(type) {
		case *dns.NS:
			records = append(records, strings.TrimRight(t.Ns, "."))
		case *dns.MX:
			records = append(records, strings.TrimRight(t.Mx, "."))
		case *dns.TXT:
			records = append(records, t.Txt...)
		}
	}
	return records, nil
}
