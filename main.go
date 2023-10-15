package main

import (
	"flag"
	"fmt"
	"net"
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

// IPInfoAPIURL is and API URL
const IPInfoAPIURL = "https://ipinfo.io/"

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

// ipsToStrings converts a slice of net.IP to a slice of string.
func ipsToStrings(ips []net.IP) []string {
	var strs []string
	for _, ip := range ips {
		strs = append(strs, ip.String())
	}
	return strs
}

// getTXT fetches the TXT records for a given domain.
func getTXT(domain string) ([]string, error) {
	return QueryDNS(domain, dns.TypeTXT, "8.8.8.8:53")
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

// fetchAPIToken fetches the IPInfo API token from environment variable or user input.
func fetchAPIToken(apiTokenFlag string) string {
	apiToken := os.Getenv("IPINFO_API_TOKEN")

	if apiTokenFlag != "" {
		apiToken = apiTokenFlag
	}

	if apiToken == "" {
		fmt.Print("IPINFO_API_TOKEN environment variable is not set.\nPlease enter your IPInfo API token: ")
		fmt.Scanln(&apiToken)
	}

	return apiToken
}

func main() {
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

	apiToken := fetchAPIToken(apiTokenFlag)

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
		promptRunner(isIP, isCIDR, input, apiToken)
	}
}

func promptRunner(isIP bool, isCIDR bool, input string, apiToken string) {
	if !isIP {
		resolvedIPPrompt(input)

		GetNSPrompt(input)

		GetMXPrompt(input)

		getTXTPrompt(input)

		getDMARCPrompt(input)

		getSPFPrompt(input)
	}

	getPTRPrompt(input)

	if isIP || isCIDR {
		asnInfo, err := GetASNInfo(input, apiToken)
		HandleResponse(asnInfo, err)
	}
}

func resolvedIPPrompt(input string) {
	ips, _ := net.LookupIP(input)
	if len(ips) > 0 {
		fmt.Printf("\033[38;5;39m Resolved IPs: \033[38;5;78m%s\033[0m\n", strings.Join(ipsToStrings(ips), ", "))
	}
}

func getTXTPrompt(input string) {
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
}

func getDMARCPrompt(input string) {
	dmarc, _ := getDMARC(input)
	if dmarc != "" {
		formattedDMARC := formatLongText(dmarc, 80, " ")
		fmt.Printf("\033[38;5;39m DMARC Record:\033[0m\n")
		fmt.Printf("\033[38;5;78m %s\033[0m\n", formattedDMARC)
	} else {
		fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88mNone\033[0m\n")
	}
}

func getPTRPrompt(input string) {
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
}
