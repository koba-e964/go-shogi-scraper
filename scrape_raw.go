package main

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type JunniScrapingRawResult struct {
	URL           string
	Name          string
	HashAlgorithm string
	Hash          []byte
	Table         [][][]string
}

// Ref: https://qiita.com/ichi_zamurai/items/91fc8bbd7dfdf7f0447f
func scrapeJunniRaw(url string) (*JunniScrapingRawResult, error) {
	scrapingResult := JunniScrapingRawResult{}
	scrapingResult.URL = url
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buffer, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	sha256sum := sha256.Sum256(buffer)
	scrapingResult.HashAlgorithm = "sha256"
	scrapingResult.Hash = sha256sum[:]

	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)

	bufferReader := bytes.NewReader(buffer)
	reader, err := charset.NewReaderLabel(detectResult.Charset, bufferReader)
	if err != nil {
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	name := document.Find("h1.ttl01").Text()
	scrapingResult.Name = name

	document.Find("tbody").Each(func(_ int, s *goquery.Selection) {
		if len(s.Children().Nodes) >= 2 {
			s.Find("tr").Each(func(x int, s *goquery.Selection) {
				data := [][]string{}
				s.Find("td").Each(func(x int, s *goquery.Selection) {
					currentCell := []string{}
					if s.HasClass("rankdown1") {
						currentCell = append(currentCell, "rankdown1")
					}
					if s.HasClass("rankdown2") {
						currentCell = append(currentCell, "rankdown2")
					}
					html, err := s.Html()
					if err != nil {
						panic(err)
					}
					currentCell = append(currentCell, strings.Split(html, "<br/>")...)
					data = append(data, currentCell)
				})
				scrapingResult.Table = append(scrapingResult.Table, data)
			})
		}
	})
	return &scrapingResult, nil
}
