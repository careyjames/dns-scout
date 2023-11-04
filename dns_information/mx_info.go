package dnsinformation

import (
	"fmt"
	"strings"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
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
		fmt.Printf(color.Blue(" MX    ✅:") + color.Grey(" ") + color.Grey(strings.Join(mx, ", ")) + constants.Newline)
	} else {
		fmt.Printf(color.Blue(" MX    ❌:") + color.Red(" None") + constants.Newline)
	}
}
