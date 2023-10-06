package main

import (
	"fmt"
	"net"

	"github.com/fatih/color"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func getRegistrar(domain string) string {
	// Fetch the WHOIS data
	result, err := whois.Whois(domain)
	if err != nil {
		return "Unknown"
	}

	// Parse the WHOIS data
	parsed, err := whoisparser.Parse(result)
	if err != nil {
		return "Unknown"
	}

	// Return the registrar name
	return parsed.Registrar.Name
}

func getNS(domain string) ([]string, error) {
	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		return nil, err
	}
	var nsNames []string
	for _, ns := range nsRecords {
		nsNames = append(nsNames, ns.Host)
	}
	return nsNames, nil
}

func getMX(domain string) ([]string, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	var mxNames []string
	for _, mx := range mxRecords {
		mxNames = append(mxNames, mx.Host)
	}
	return mxNames, nil
}

func getTXT(domain string) ([]string, error) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err
	}
	return txtRecords, nil
}

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

func getPTR(domain string) ([]string, error) {
	ptrRecords, err := net.LookupAddr(domain)
	if err != nil {
		return nil, err
	}
	return ptrRecords, nil
}

func main() {
	for {
		var domain string
		color.Cyan("Enter domain (or 'exit' to quit): ")
		fmt.Scanln(&domain)

		if domain == "exit" {
			return
		}
		registrar := getRegistrar(domain)
		color.Cyan("Registrar: %s", registrar)

		ns, _ := getNS(domain)
		mx, _ := getMX(domain)
		txt, _ := getTXT(domain)
		dmarc, _ := getDMARC(domain)
		ptr, _ := getPTR(domain)

		color.Cyan("Name Servers: %v", ns)
		color.Cyan("MX Records: %v", mx)
		color.Cyan("TXT Records: %v", txt)
		color.Cyan("DMARC Record: %v", dmarc)
		color.Cyan("PTR Records: %v", ptr)
	}
}
