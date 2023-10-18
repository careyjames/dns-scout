package dnsinformation

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/careyjames/DNS-Scout/color"
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
	return &ipInfo, nil
}

// GetASNInfoPrompt handles response for asn info
func GetASNInfoPrompt(input string, apiToken string) {
	asnInfo, err := GetASNInfo(input, apiToken)
	if err == nil {
		fmt.Printf(color.Blue(" ASN Information: ") + constants.Newline)
		fmt.Printf(color.Blue(" IP: ") + constants.Newline)
		fmt.Printf(color.Blue(" Domain: ") + color.Green(asnInfo.Domain) + constants.Newline)
		fmt.Printf(color.Blue(" HostName: ") + color.Green(asnInfo.Hostname) + constants.Newline)
		fmt.Printf(color.Blue(" City: ") + color.Green(asnInfo.City) + constants.Newline)
		fmt.Printf(color.Blue(" Region: ") + color.Green(asnInfo.Region) + constants.Newline)
		fmt.Printf(color.Blue(" Country: ") + color.Green(asnInfo.Country) + constants.Newline)
		fmt.Printf(color.Blue(" Location: ") + color.Green(asnInfo.Loc) + constants.Newline)
		fmt.Printf(color.Blue(" Organization: ") + color.Green(asnInfo.Org) + constants.Newline)
		fmt.Printf(color.Blue(" Postal Code: ") + color.Green(asnInfo.Postal) + constants.Newline)
		fmt.Printf(color.Blue(" Timezone: ") + color.Green(asnInfo.Timezone) + constants.Newline)
		asnInfoStrs := []string{}
		for k, v := range asnInfo.ASN {
			asnInfoStrs = append(asnInfoStrs, fmt.Sprintf("%s: %v", k, v))
		}
		fmt.Printf(color.Blue(" ASN: ") + color.Green(strings.Join(asnInfoStrs, ", ")) + constants.Newline)
	} else {
		fmt.Println(" Error fetching ASN Information:", err)
	}
}
