package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	clients "github.com/careyjames/DNS-Scout/clients"
	constants "github.com/careyjames/DNS-Scout/constant"
	dnsinformation "github.com/careyjames/DNS-Scout/dns_information"
	color "github.com/fatih/color"

	"github.com/briandowns/spinner"
	"github.com/chzyer/readline"
)

// getDMARC fetches the DMARC record for a given domain.
func getDMARC(domain string) (string, error) {
	txtRecords, err := dnsinformation.GetTXT("_dmarc." + domain)
	if len(txtRecords) <= 0 {
		txtRecords, _ = dnsinformation.GetDMARCRecordNSLookup(domain)
	}
	if err != nil {
		return "", err
	}
	for _, record := range txtRecords {
		if len(record) > 8 && record[:8] == "v=DMARC1" {
			return record, nil
		} else if record[:6] == "vDMARC" {
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

	if !valid && strings.Contains(strings.ToLower(record), "spf") {
		colorCode = "\033[38;5;88m" // Green for invalid SPF
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

func main() {
	// Check "version" argument
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "version" {
		fmt.Println("DNS-Scout version:", constants.Version)
		return
	}
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

	apiToken := clients.FetchAPIToken(apiTokenFlag)

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

		s.Stop() // Stop the spinner

		promptRunner(isIP, isCIDR, input, apiToken)
	}
}

func promptRunner(isIP bool, isCIDR bool, input string, apiToken string) {
	if !isIP {
		dnsinformation.GetRegistrarPromt(input, isIP)

		dnsinformation.ResolvedIPPrompt(input)

		dnsinformation.GetNSPrompt(input)

		dnsinformation.GetMXPrompt(input)

		dnsinformation.GetTXTPrompt(input)

		getDMARCPrompt(input)

		dnsinformation.GetSPFPrompt(input)
	}

	getPTRPrompt(input, isIP)

	if isIP || isCIDR {
		asnInfo, err := GetASNInfo(input, apiToken)
		HandleResponse(asnInfo, err)
	}
}

func isValidDMARC(record string) bool {
	if len(record) > 8 && record[:8] == "v=DMARC1" {
		return true
	}
	return false
}

func getDMARCPrompt(input string) {
	dmarc, _ := getDMARC(input)
	if dmarc != "" {
		if isValidDMARC(dmarc) {
			formattedDMARC := formatLongText(dmarc, 80, " ")
			fmt.Printf("\033[38;5;39m DMARC Record:\033[0m\n")
			fmt.Printf("\033[38;5;78m %s\033[0m\n", formattedDMARC)
		} else {
			fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88m%s\033[0m\n", dmarc[8:])
		}
	} else {
		fmt.Printf("\033[38;5;39m DMARC Record: \033[0m\033[38;5;88mNone\033[0m\n")
	}
}

func getPTRPrompt(input string, isIp bool) {
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
		if !isIp {
			fmt.Printf("\033[38;5;39m PTR Records: \033[38;5;78m%s\033[0m\n", ptrStr)
		} else {
			fmt.Printf("\033[38;5;39m PTR Records: \033[38;5;78m%s\033[0m", ptrStr)
		}
	} else {
		if !isIp {
			fmt.Printf("\033[38;5;39m PTR Records: \033[0m\033[38;5;222mNone\033[0m\n")
		} else {
			fmt.Printf("\033[38;5;39m PTR Records: \033[0m\033[38;5;222mNone\033[0m")
		}
	}
}
