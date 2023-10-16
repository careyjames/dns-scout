package main

import (
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

// getMX fetches the MX records for a given domain.
func getMX(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeMX, "8.8.8.8:53")
}

// GetMXPrompt is MX prompt
func GetMXPrompt(input string) {
	mx, _ := getMX(input)
	if len(mx) > 1 || len(mx) == 1 && mx[0] != "" {
		fmt.Printf("\033[38;5;39m MX Records: \033[38;5;78m%s\033[0m\n", strings.Join(mx, ", "))
	} else {
		fmt.Printf("\033[38;5;39m MX Records: \033[0m\033[38;5;88mNo MX, No email.\033[0m\n")
	}
}
