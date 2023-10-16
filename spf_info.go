package main

import (
	"fmt"
	"strings"
)

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

// GetSPFPrompt is prompt for spf
func getSPFPrompt(input string) {
	spfValid, spfRecord, _ := getSPF(input)
	txt, _ := getTXT(input)
	countTxt := 0
	if len(txt) > 1 {
		for _, record := range txt {
			if strings.Contains(strings.ToLower(record), "spf") {
				countTxt = countTxt + 1
			}
		}
	}
	if countTxt > 1 {
		fmt.Printf("\033[38;5;39m TXT Records: \033[0m\033[38;5;88mCan't have two SPF!\033[0m\n")
	}

	if spfValid || spfRecord != "No SPF record" {
		coloredSPFRecord := colorCodeSPFRecord(spfRecord, spfValid)
		fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
	} else {
		coloredSPFRecord := colorCodeSPFRecord(spfRecord, false) // "No SPF record" will be red
		fmt.Printf("\033[38;5;39m SPF Record: %s\033[0m\n", coloredSPFRecord)
	}
}
