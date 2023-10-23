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

func getURL() string {
	return constants.IPInfoAPIURL
}

// GetASNInfo fetches ASN information for a given IP address.
func GetASNInfo(ip string, apiToken string) (*dto.IPInfoResponse, error) {
	// Removed apiToken from here as it's now passed as an argument
	resp, err := http.Get(getURL() + ip + "?token=" + apiToken)
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
		// Check if ASN information is empty or null
		if len(asnInfo.ASN) == 0 {
			fmt.Printf(color.Blue(" ASN   ❌: ") + color.Red("None") + constants.Newline)
		} else {
			fmt.Printf(color.Blue("") + constants.Newline)
			fmt.Printf(color.Blue(" IP: ") + color.Grey(asnInfo.IP) + constants.Newline)
			fmt.Printf(color.Blue(" Domain: ") + color.Grey(asnInfo.Domain) + constants.Newline)
			fmt.Printf(color.Blue(" HostName: ") + color.Grey(asnInfo.Hostname) + constants.Newline)
			fmt.Printf(color.Blue(" City: ") + color.Grey(asnInfo.City) + constants.Newline)
			fmt.Printf(color.Blue(" Region: ") + color.Grey(asnInfo.Region) + constants.Newline)
			fmt.Printf(color.Blue(" Country: ") + color.Grey(asnInfo.Country) + constants.Newline)
			fmt.Printf(color.Blue(" Location: ") + color.Grey(asnInfo.Loc) + constants.Newline)
			fmt.Printf(color.Blue(" Organization: ") + color.Grey(asnInfo.Org) + constants.Newline)
			fmt.Printf(color.Blue(" Postal Code: ") + color.Grey(asnInfo.Postal) + constants.Newline)
			fmt.Printf(color.Blue(" Timezone: ") + color.Grey(asnInfo.Timezone) + constants.Newline)
			asnInfoStrs := []string{}
			for k, v := range asnInfo.ASN {
				asnInfoStrs = append(asnInfoStrs, fmt.Sprintf("%s: %v", k, v))
			}
			fmt.Printf(color.Blue(" ASN   ✅: ") + color.Grey(strings.Join(asnInfoStrs, ", ")) + constants.Newline)
		}
	} else {
		fmt.Println(" Error fetching ASN Information:", err)
	}
}
