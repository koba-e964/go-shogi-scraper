package main

import (
	"fmt"
)

func main() {
	url := "https://www.shogi.or.jp/match/junni/2023/82a/index.html"

	scrapingResult, err := scrapeJunniRaw(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(scrapingResult.Name)
	fmt.Printf("%s: %x\n", scrapingResult.HashAlgorithm, scrapingResult.Hash)
}
