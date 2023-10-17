package dnsinformation

import (
	"fmt"
	"net"
	"strings"
)

// IpsToStrings converts a slice of net.IP to a slice of string.
func ipsToStrings(ips []net.IP) []string {
	var strs []string
	for _, ip := range ips {
		strs = append(strs, ip.String())
	}
	return strs
}

func ResolvedIPPrompt(input string) {
	ips, _ := net.LookupIP(input)
	if len(ips) > 0 {
		fmt.Printf("\033[38;5;39m Resolved IPs: \033[38;5;78m%s\033[0m\n", strings.Join(ipsToStrings(ips), ", "))
	}
}
