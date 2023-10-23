package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
	"github.com/miekg/dns"
)

// getMX fetches the MX records for a given domain.
func getMX(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeMX, constants.GooglePublicDNS)
}

// GetMXPrompt is MX prompt
func GetMXPrompt(input string) {
	mx, _ := getMX(input)
	if len(mx) > 1 || len(mx) == 1 && mx[0] != "" {
		fmt.Printf(color.Blue(" MX Records: ") + color.Green("✅ ") + color.Green(strings.Join(mx, ", ")) + constants.Newline)
	} else {
		fmt.Printf(color.Blue(" MX Records: ") + color.Red("❌ None") + constants.Newline)
	}
}
