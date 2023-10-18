package dnsinformation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	constants "github.com/careyjames/DNS-Scout/constant"
	"github.com/careyjames/DNS-Scout/dto"
)

// GetASNInfo fetches ASN information for a given IP address.
func GetASNInfo(ip string, apiToken string) (*dto.IPInfoResponse, error) {
	// Removed apiToken from here as it's now passed as an argument
	resp, err := http.Get(constants.IPInfoAPIURL + ip + "?token=" + apiToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo dto.IPInfoResponse
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return nil, err
	}
	// Add this line for debugging
	//fmt.Printf("Debug: Received IPInfo API response: %+v\n", ipInfo)
	return &ipInfo, nil
}

// GetASNInfoPrompt handles response for asn info
func GetASNInfoPrompt(input string, apiToken string) {
	asnInfo, err := GetASNInfo(input, apiToken)
	if err == nil {
		fmt.Printf("\033[38;5;39m\n ASN Information: \n\033[0m")
		fmt.Printf("\033[38;5;39m IP: \033[38;5;78m%s\033[0m\n", asnInfo.IP)
		fmt.Printf("\033[38;5;39m Domain: \033[38;5;78m%s\033[0m\n", asnInfo.Domain)
		fmt.Printf("\033[38;5;39m Hostname: \033[38;5;78m%s\033[0m\n", asnInfo.Hostname)
		fmt.Printf("\033[38;5;39m City: \033[38;5;78m%s\033[0m\n", asnInfo.City)
		fmt.Printf("\033[38;5;39m Region: \033[38;5;78m%s\033[0m\n", asnInfo.Region)
		fmt.Printf("\033[38;5;39m Country: \033[38;5;78m%s\033[0m\n", asnInfo.Country)
		fmt.Printf("\033[38;5;39m Location: \033[38;5;78m%s\033[0m\n", asnInfo.Loc)
		fmt.Printf("\033[38;5;39m Organization: \033[38;5;78m%s\033[0m\n", asnInfo.Org)
		fmt.Printf("\033[38;5;39m Postal Code: \033[38;5;78m%s\033[0m\n", asnInfo.Postal)
		fmt.Printf("\033[38;5;39m Timezone: \033[38;5;78m%s\033[0m\n", asnInfo.Timezone)
		asnInfoStrs := []string{}
		for k, v := range asnInfo.ASN {
			asnInfoStrs = append(asnInfoStrs, fmt.Sprintf("%s: %v", k, v))
		}
		fmt.Printf("\033[38;5;39m ASN: \033[38;5;78m%s\033[0m\n", strings.Join(asnInfoStrs, ", "))

	} else {
		fmt.Println(" Error fetching ASN Information:", err)
	}
}
