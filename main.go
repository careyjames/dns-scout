// The main function in this code is a command-line tool that fetches and displays various information
// about a domain or IP address, including registrar, DNS records, SPF and DMARC records, PTR records,
// and ASN information.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	color "github.com/fatih/color"

	"github.com/briandowns/spinner"
	"github.com/chzyer/readline"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/miekg/dns"
)

const IPInfoAPIURL = "https://ipinfo.io/"

// IPInfoResponse struct holds the response from the IPInfo API
type IPInfoResponse struct {
	IP       string `json:"ip"`
	Domain   string `json:"domain"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

// getRegistrar fetches the registrar information for a given domain.
func getRegistrar(domain string) string {
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

// getASNInfo fetches ASN information for a given IP address.
func getASNInfo(ip string, apiToken string) (*IPInfoResponse, error) {
	// Removed apiToken from here as it's now passed as an argument
	resp, err := http.Get(IPInfoAPIURL + ip + "?token=" + apiToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo IPInfoResponse
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		return nil, err
	}
	// Add this line for debugging
	//fmt.Printf("Debug: Received IPInfo API response: %+v\n", ipInfo)
	return &ipInfo, nil
}

// queryDNS performs a DNS query for a given domain and DNS record type.
func queryDNS(domain string, dnsType uint16, server string) ([]string, error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dnsType)

	r, _, err := c.Exchange(m, server)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, ans := range r.Answer {
		switch t := ans.(type) {
		case *dns.NS:
			records = append(records, strings.TrimRight(t.Ns, "."))
		case *dns.MX:
			records = append(records, strings.TrimRight(t.Mx, "."))
		case *dns.TXT:
			records = append(records, t.Txt...)
		}
	}
	return records, nil
}

// ipsToStrings converts a slice of net.IP to a slice of string.
func ipsToStrings(ips []net.IP) []string {
	var strs []string
	for _, ip := range ips {
		strs = append(strs, ip.String())
	}
	return strs
}

// getNS fetches the NS records for a given domain.
func getNS(domain string) ([]string, error) {
	googleRecords, err1 := queryDNS(domain, dns.TypeNS, "8.8.8.8:53")
	cloudflareRecords, err2 := queryDNS(domain, dns.TypeNS, "1.1.1.1:53")

	if err1 != nil && err2 != nil {
		return nil, fmt.Errorf("both DNS queries failed")
	}

	// Merge and deduplicate records
	recordMap := make(map[string]bool)
	for _, record := range googleRecords {
		recordMap[record] = true
	}
	for _, record := range cloudflareRecords {
		recordMap[record] = true
	}

	var mergedRecords []string
	for record := range recordMap {
		mergedRecords = append(mergedRecords, record)
	}

	return mergedRecords, nil
}

// getMX fetches the MX records for a given domain.
func getMX(domain string) ([]string, error) {
	return queryDNS(domain, dns.TypeMX, "8.8.8.8:53")
}

// getTXT fetches the TXT records for a given domain.
func getTXT(domain string) ([]string, error) {
	return queryDNS(domain, dns.TypeTXT, "8.8.8.8:53")
}

// getSPF fetches and analyzes the SPF record for a given domain.
func getSPF(domain string) (bool, string, string) {
	txtRecords, err := getTXT(domain)
	if err != nil {
		return false, "Error fetching TXT records", ""
	}
	for _, record := range txtRecords {
		suffix := ""
		if strings.Contains(record, "-all") {
			suffix = "-all"
		} else if strings.Contains(record, "~all") {
			suffix = "~all"
		}

		if strings.HasPrefix(record, "v=spf1") {
			return true, record, suffix
		} else if strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
			return false, record, suffix
		}
	}
	return false, " No SPF record", ""
}

// getDMARC fetches the DMARC record for a given domain.
func getDMARC(domain string) (string, error) {
	txtRecords, err := getTXT("_dmarc." + domain)
	if err != nil {
		return "", err
	}
	for _, record := range txtRecords {
		if len(record) > 8 && record[:8] == "v=DMARC1" {
			return record, nil
		}
	}
	return "", nil
}

// formatLongText formats long text strings for better readability.
func formatLongText(text string, threshold int, indent string) string {
	if len(text) <= threshold {
		return text
	}

	var result strings.Builder
	for len(text) > threshold {
		splitAt := strings.LastIndex(text[:threshold], " ")
		if splitAt == -1 {
			splitAt = threshold
		}
		result.WriteString(text[:splitAt] + "\n" + indent)
		text = text[splitAt+1:]
	}
	result.WriteString(text)
	return result.String()
}

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

func colorCodeSPFRecord(record string, valid bool) string {
	colorCode := "\033[38;5;78m" // Default to green for non-SPF records

	if strings.HasPrefix(record, "v=spf1") || strings.Contains(record, "spf") || strings.Contains(record, "-all") || strings.Contains(record, "~all") {
		colorCode = "\033[38;5;88m" // Default to red for malformed or misspelled SPF
		if valid {
			colorCode = "\033[38;5;78m" // Green for valid SPF
		}
	}

	if record == " No SPF record" {
		colorCode = "\033[38;5;88m" // Red for "No SPF record"
	}

	if strings.Contains(record, "-all") {
		record = strings.ReplaceAll(record, "-all", "\033[38;5;222m-all\033[0m")
	} else if strings.Contains(record, "~all") {
		record = strings.ReplaceAll(record, "~all", "\033[38;5;78m~all\033[0m")
	}

	return fmt.Sprintf("%s%s\033[0m", colorCode, record)
}

// main is the entry point of the application.
func main() {
	// Add these lines right here, at the beginning of the main function
	var apiTokenFlag string
	flag.StringVar(&apiTokenFlag, "api-token", "", "IPInfo API token")
	flag.Parse()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:              " \033[38;5;39m🌎\033[38;5;39m ",
		HistoryFile:         ".tmp-history",
		AutoComplete:        nil,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: nil,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	// Read API token from environment variable
	apiToken := os.Getenv("IPINFO_API_TOKEN")

	// Override with command-line argument if provided
	if apiTokenFlag != "" {
		apiToken = apiTokenFlag
	}

	// If API token is still not set, prompt the user
	if apiToken == "" {
		fmt.Print("IPINFO_API_TOKEN environment variable is not set.\nPlease enter your IPInfo API token: ")
		fmt.Scanln(&apiToken)
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Use the dots character set and update every 100ms

	for {
		color.New(color.FgHiWhite).Println(" Enter domain, IP (or 'exit' to quit): ")
		fmt.Println("\033[38;5;39m ------------------------------------\033[0m")

		input, err := rl.Readline()

		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}
		if input == "exit" {
			return
		}

		s.Start() // Start the spinner

		isIP := net.ParseIP(input) != nil
		_, _, err = net.ParseCIDR(input)
		isCIDR := err == nil

		registrar := getRegistrar(input)
		s.Stop() // Stop the spinner

		if !isIP || (isIP && registrar != "Unknown or Classified") {
			fmt.Printf("\033[38;5;39m Registrar: \033[38;5;78m%s\033[0m\n", registrar)
		}

		if !isIP {
			ips, _ := net.LookupIP(input)
			if len(ips) > 0 {
				fmt.Printf("\033[38;5;39m Resolved IPs: \033[38;5;78m%s\033[0m\n", strings.Join(ipsToStrings(ips), ", "))
			}

			ns, _ := getNS(input)
			if len(ns) > 0 {
				fmt.Printf("\033[38;5;39m Name Servers: \033[38;5;78m%s\033[0m\n", strings.Join(ns, ", "))
			} else {
				fmt.Printf("\033[38;5;39m Name Servers: \033[0m\033[38;5;88mNone\033[0m\n")
			}

			mx, _ := getMX(input)
			if len(mx) > 0 {
				fmt.Printf("\033[38;5;39m MX Records: \033[38;5;78m%s\033[0m\n", strings.Join(mx, ", "))
			} else {
				fmt.Printf("\033[38;5;39m MX Records: \033[0m\033[38;5;88mNo MX, No email.\033[0m\n")
			}

			txt, _ := getTXT(input)
			if len(txt) > 0 {
				fmt.Printf("\033[38;5;39m TXT Records:\033[0m\n")
				for _, record := range txt {
					isValidSPF := strings.HasPrefix(record, "v=spf1")
					coloredRecord := colorCodeSPFRecord(record, isValidSPF)
					fmt.Printf(" %s\n", coloredRecord)
				}
			} else {
				fmt.Printf("\033[38;5;39m TXT Records: \033[0m\033[38;5;88mNone\033[0m\n")
			}

			dmarc, _ := getDMARC(input)
			if dmarc != "" {
				formattedDMARC := formatLongText(dmarc, 80, " ")
				fmt.Printf("\033[38;5;39m DMARC Record:\033[0m\n")
				fmt.Printf("\033[38;5;78m %s\033[0m\n", formattedDMARC)
			} else {
				fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88mNone\033[0m\n")
			}

			spfValid, spfRecord, _ := getSPF(input)

			if spfValid || spfRecord != "No SPF record" {
				coloredSPFRecord := colorCodeSPFRecord(spfRecord, spfValid)
				fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
			} else {
				coloredSPFRecord := colorCodeSPFRecord(spfRecord, false) // "No SPF record" will be red
				fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
			}
		}

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
			fmt.Printf("\033[38;5;39m PTR Records: \033[38;5;78m%s\033[0m\n", ptrStr)
		} else {
			fmt.Printf("\033[38;5;39m PTR Records: \033[0m\033[38;5;222mNone\033[0m\n")
		}

		if isIP || isCIDR {
			asnInfo, err := getASNInfo(input, apiToken)
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
			} else {
				fmt.Println(" Error fetching ASN Information:", err)
			}
		}
	}
}
