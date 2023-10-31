package dnsinformation

import (
	"fmt"
	"net"
	"strings"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
)

// getPTR fetches the PTR records for a given domain.
func getPTR(domain string) ([]string, error) {
	// mx - ptr (AND OPERATION)
	// TRUE - TRUE : TRUE
	// FALSE - TRUE : FALSE
	// TRUE - PTR : FALSE
	// FALSE - FALSE : FALSE

	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ptrRecords []string
	for _, ip := range ips {
		ptrs, err := net.LookupAddr(ip.String())
		if err != nil {
			continue
		}
		ptrRecords = append(ptrRecords, ptrs...)
	}
	return ptrRecords, nil
}

// GetPTRPrompt prompt
func GetPTRPrompt(input string, isIp bool) {
	ptr, _ := getPTR(input)
	if len(ptr) > 0 {
		// Remove the trailing period from each PTR record
		for i, record := range ptr {
			if strings.HasSuffix(record, ".") {
				ptr[i] = record[:len(record)-1]
			}
		}
		// Join the records with a comma and a space, then replace ", " at the end of each line with a line break followed by a space
		ptrStr := strings.Join(ptr, ", ")
		ptrStr = strings.ReplaceAll(ptrStr, ", ", ",\n ")
		if !isIp && !isMx(input) {
			fmt.Printf(color.Blue(" PTR   ✅: ") + color.Grey(ptrStr) + constants.Newline)
		} else {
			// we might need to delete this
			fmt.Printf(color.Blue(" PTR   ✅: ") + color.Grey(ptrStr))
		}
	} else {
		if !isIp {
		        // FIXME: add mapping for google microsoft for their service
			fmt.Printf(color.Blue(" PTR   🟢: ") + color.Grey("None, Google and Microsoft 365 use shared IPs, this is ok.") + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" PTR   ❌: ") + color.Red("None"))
		}
	}
}

func isMx(input string) bool {
	x, _ := getMX(input)
	return len(x) > 0
}
