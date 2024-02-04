package players

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	scraper "github.com/koba-e964/go-shogi-scraper"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

const DefaultURL = "https://www.shogi.or.jp/player/japanese_syllabary_order.html"

var playerInfoURL = regexp.MustCompile(`^/player/(.*)/(\d+).html$`)

type PlayerList struct {
	URL           string           `json:"url"`
	RetrievalTime string           `json:"retrieval_time"`
	HashAlgorithm string           `json:"hash_algorithm"`
	Hash          []byte           `json:"hash"`
	Players       []scraper.Player `json:"players"`
}

func (p *PlayerList) GetURL() string {
	return p.URL
}

func (p *PlayerList) GetRetrievalTime() string {
	return p.RetrievalTime
}

type players []scraper.Player

func (p players) Len() int      { return len(p) }
func (p players) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p players) Less(i, j int) bool {
	if p[i].Kind != p[j].Kind {
		return p[i].Kind < p[j].Kind
	}
	return p[i].ID < p[j].ID
}

func ScrapePlayerList(url string) (*PlayerList, error) {
	scrapingResult := PlayerList{}
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

	scrapingResult.Players = make([]scraper.Player, 0)
	document.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			// ignore
			return
		}
		name := s.Text()
		if playerInfoURL.MatchString(href) {
			kind := playerInfoURL.FindStringSubmatch(href)[1]
			id, err := strconv.Atoi(playerInfoURL.FindStringSubmatch(href)[2])
			if err != nil {
				panic(err) // should not happen
			}
			player := scraper.Player{Kind: kind, ID: id, Name: name}
			scrapingResult.Players = append(scrapingResult.Players, player)
		}
	})
	sort.Sort(players(scrapingResult.Players))
	return &scrapingResult, nil
}
