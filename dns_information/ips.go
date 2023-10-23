package dnsinformation

import (
	"fmt"
	"net"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

// IpsToStrings converts a slice of net.IP to a slice of string.
func ipsToStrings(ips []net.IP) []string {
	var strs []string
	for _, ip := range ips {
		strs = append(strs, ip.String())
	}
	return strs
}

// ResolvedIPPrompt...
func ResolvedIPPrompt(input string) {
	ips, _ := net.LookupIP(input)
	if len(ips) > 0 {
		fmt.Printf(color.Blue(" IPs   🟢: ") + color.Grey(strings.Join(ipsToStrings(ips), ", ")) + constants.Newline)
	} else {
		fmt.Printf(color.Blue(" IPs   ❌: ") + color.Red("None") + constants.Newline)
	}
}
