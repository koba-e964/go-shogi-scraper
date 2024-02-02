package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s URL\n", os.Args[0])
		return
	}
	url := os.Args[1]

	scrapingRawResult, err := ScrapeJunniRaw(url)
	if err != nil {
		panic(err)
	}
	scrapingResult, err := ParseRawJunni(scrapingRawResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(scrapingRawResult.Name)
	fmt.Printf("%s: %x\n", scrapingRawResult.HashAlgorithm, scrapingRawResult.Hash)
	// write as JSON to file
	rawJSON, err := json.MarshalIndent(scrapingRawResult, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("scraping_raw_result.json", rawJSON, 0o644)
	resultJSON, err := json.MarshalIndent(scrapingResult, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("scraping_result.json", resultJSON, 0o644)
}
