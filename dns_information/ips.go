package dnsinformation

import (
	"fmt"
	"net"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
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
		fmt.Printf(color.Blue(" Resolved IPs: ") + color.Green(strings.Join(ipsToStrings(ips), ", ")))
	}
}
