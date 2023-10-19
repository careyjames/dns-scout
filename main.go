package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	clients "github.com/careyjames/DNS-Scout/clients"
	constants "github.com/careyjames/DNS-Scout/constant"
	dnsinformation "github.com/careyjames/DNS-Scout/dns_information"
	color "github.com/fatih/color"

	"github.com/briandowns/spinner"
	"github.com/chzyer/readline"
)

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
		dnsinformation.GetDMARCPrompt(input)
		dnsinformation.GetDKIMPrompt(input)
		dnsinformation.GetSPFPrompt(input)
	}
	dnsinformation.GetPTRPrompt(input, isIP)
	if isIP || isCIDR {
		dnsinformation.GetASNInfoPrompt(input, apiToken)
	}
}
