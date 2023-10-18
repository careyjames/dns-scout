package dto

// IPInfoResponse struct holds the response from the IPInfo API
type IPInfoResponse struct {
	ASN      map[string]interface{} `json:"asn"`
	IP       string                 `json:"ip"`
	Domain   string                 `json:"domain"`
	Hostname string                 `json:"hostname"`
	City     string                 `json:"city"`
	Region   string                 `json:"region"`
	Country  string                 `json:"country"`
	Loc      string                 `json:"loc"`
	Org      string                 `json:"org"`
	Postal   string                 `json:"postal"`
	Timezone string                 `json:"timezone"`
	Readme   string                 `json:"readme"`
}
