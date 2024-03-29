package clients

import (
	"fmt"
	"os"
)

const (
	IPINFO_API_TOKEN = "IPINFO_API_TOKEN"
)

// FetchAPIToken fetches the IPInfo API token from environment variable or user input.
func FetchAPIToken(apiTokenFlag string) string {
	apiToken := os.Getenv(IPINFO_API_TOKEN)
	if apiTokenFlag != "" {
		apiToken = apiTokenFlag
	}
	if apiToken == "" {
		fmt.Print("IPINFO_API_TOKEN environment variable not set.\nEnter ipinfo.io API token, or press return to skip ASN: ")
		fmt.Scanln(&apiToken)
	}
	return apiToken
}
