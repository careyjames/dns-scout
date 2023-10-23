package dnsinformation

import (
	"fmt"
	"net"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
)

// getPTR fetches the PTR records for a given domain.
func getPTR(domain string) ([]string, error) {
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
		if !isIp {
			fmt.Printf(color.Blue(" PTR Records: ✅ ") + color.Green(ptrStr) + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" PTR Records: ✅ ") + color.Green(ptrStr))
		}
	} else {
		if !isIp {
			fmt.Printf(color.Blue(" PTR Records: ✅ ") + color.Green("None, Google and Microsoft 365 use shared IPs, this is ok.") + constants.Newline)
		} else {
			fmt.Printf(color.Blue(" PTR Records: ❌ ") + color.Red("None"))
		}
	}
}
