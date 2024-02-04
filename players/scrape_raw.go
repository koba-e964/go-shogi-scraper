package players

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	scraper "github.com/koba-e964/go-shogi-scraper"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type PlayerInfoRaw struct {
	ID            int    `json:"id"`
	NameJP        string `json:"name_jp"`
	NameEN        string `json:"name_en"`
	URL           string `json:"url"`
	RetrievalTime string `json:"retrieval_time"`
	HashAlgorithm string `json:"hash_algorithm"`
	Hash          []byte `json:"hash"`
	Birthday      string `json:"birthday"`
	Birthplace    string `json:"birthplace"`
	MentorName    string `json:"mentor_name"`
}

func (p *PlayerInfoRaw) GetURL() string {
	return p.URL
}

func (p *PlayerInfoRaw) GetRetrievalTime() string {
	return p.RetrievalTime
}

func ScrapePlayerInfoRaw(url string, player *scraper.Player, playersList *PlayerList) (*PlayerInfoRaw, error) {
	scrapingResult := PlayerInfoRaw{}

	scrapingResult.ID = player.ID
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

	scrapingResult.RetrievalTime = time.Now().UTC().Format(time.RFC3339)

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

	document.Find(".jp").Each(func(_ int, s *goquery.Selection) {
		scrapingResult.NameJP = s.Text()
	})
	document.Find(".en").Each(func(_ int, s *goquery.Selection) {
		scrapingResult.NameEN = s.Text()
	})

	document.Find("tbody").Each(func(_ int, s *goquery.Selection) {
		s.Find("tr").Each(func(_ int, s *goquery.Selection) {
			key := ""
			value := ""
			s.Find("th").Each(func(_ int, s *goquery.Selection) {
				key = s.Text()
			})
			s.Find("td").Each(func(_ int, s *goquery.Selection) {
				value = s.Text()
			})
			switch key {
			case "棋士番号":
			case "生年月日":
				scrapingResult.Birthday = value
			case "出身地":
				scrapingResult.Birthplace = value
			case "師匠":
				scrapingResult.MentorName = value
			default:
			}
		})
	})

	return &scrapingResult, nil
}
