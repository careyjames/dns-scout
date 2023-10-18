package dnsinformation

import (
	"fmt"

	"github.com/careyjames/DNS-Scout/color"
	constants "github.com/careyjames/DNS-Scout/constant"
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

func GetRegistrarPromt(input string, isIP bool) {
	registrar := getRegistrar(input)

	if registrar == ErrorMessage {
		fmt.Printf(color.Blue(" Registrar: ") + color.Green("Unknown or ") + color.Yellow("Classified") + constants.Newline)
	} else {
		if !isIP || (isIP && registrar != ErrorMessage) {
			fmt.Printf(color.Blue(" Registrar: ") + color.Green(registrar) + constants.Newline)
		}
	}
}
