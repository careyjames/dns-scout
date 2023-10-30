package dnsinformation

import (
	"fmt"

	"github.com/careyjames/dns-scout/color"
	constants "github.com/careyjames/dns-scout/constant"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

const (
	ErrorMessage = "Unknown or Classified"
)

// getRegistrar fetches the registrar information for a given domain.
func getRegistrar(domain string) string {
	result, err := whois.Whois(domain)
	if err != nil {
		return ErrorMessage
	}
	parsed, err := whoisparser.Parse(result)
	if err != nil {
		return ErrorMessage
	}
	if parsed.Registrar != nil {
		return parsed.Registrar.Name
	}
	return ErrorMessage
}

// GetRegistrarPromt ...
func GetRegistrarPromt(input string, isIP bool) {
	registrar := getRegistrar(input)

	if registrar == ErrorMessage {
		fmt.Printf(color.Blue(" Reg   🟡: ") + color.Grey("Unknown or ") + color.Yellow("Classified") + constants.Newline)
	} else {
		if !isIP || (isIP && registrar != ErrorMessage) {
			fmt.Printf(color.Blue(" Reg   🟢: ") + color.Grey(registrar) + constants.Newline)
		}
	}
}
