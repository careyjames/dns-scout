package dnsinformation

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

// GetRegistrar fetches the registrar information for a given domain.
func GetRegistrar(domain string) string {
	result, err := whois.Whois(domain)
	if err != nil {
		return "Unknown or Classified"
	}

	parsed, err := whoisparser.Parse(result)
	if err != nil {
		return "Unknown or Classified"
	}

	if parsed.Registrar != nil {
		return parsed.Registrar.Name
	}

	return "Unknown or Classified"
}
