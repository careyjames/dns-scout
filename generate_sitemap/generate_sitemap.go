package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Url struct {
	Loc string `xml:"loc"`
}

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

func GenerateSitemap() {
	urls := []Url{
		{Loc: "https://dns-scout.com/"},
		{Loc: "https://dns-scout.com/about"},
		// Add more URLs here
	}

	urlset := Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	output, err := xml.MarshalIndent(urlset, " ", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	file, err := os.Create("sitemap.xml")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(xml.Header)
	file.Write(output)
}

func main() {
	GenerateSitemap()
}
