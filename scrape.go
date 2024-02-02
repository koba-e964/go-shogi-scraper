package main

import (
	"fmt"
	"sort"
)

type nameIdentifier struct {
	names []string
	cache map[string]int
}

func isSubsequence[T comparable](subsequence, sequence []T) bool {
	subIndex := 0
	mainIndex := 0
	for subIndex < len(subsequence) && mainIndex < len(sequence) {
		if subsequence[subIndex] == sequence[mainIndex] {
			subIndex++
		}
		mainIndex++
	}
	return subIndex == len(subsequence)
}

func (n *nameIdentifier) getIndex(shortName string) int {
	if id, ok := n.cache[shortName]; ok {
		return id
	}
	for i, name := range n.names {
		// subsequence?
		if isSubsequence([]rune(shortName), []rune(name)) {
			n.cache[shortName] = i
			return i
		}
	}
	return -1
}

type Player struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type Result struct {
	Index1        int    `json:"index1"`
	Index2        int    `json:"index2"`
	Winner        int    `json:"winner,omitempty"`     // 0: not decided, 1: Index1, 2: Index2
	PlayFirst     int    `json:"play_first,omitempty"` // 0: not decided, 1: Index1, 2: Index2
	Round         int    `json:"round"`
	GameRecordURL string `json:"game_record_url,omitempty"`
}

type results []Result

func (r results) Len() int      { return len(r) }
func (r results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r results) Less(i, j int) bool {
	if r[i].Round != r[j].Round {
		return r[i].Round < r[j].Round
	}
	if r[i].Index1 != r[j].Index1 {
		return r[i].Index1 < r[j].Index1
	}
	if r[i].Index2 != r[j].Index2 {
		return r[i].Index2 < r[j].Index2
	}
	return false
}

type JunniScrapingResult struct {
	URL           string   `json:"url"`
	RetrievalTime string   `json:"retrieval_time"`
	Name          string   `json:"name"`
	HashAlgorithm string   `json:"hash_algorithm"`
	Hash          []byte   `json:"hash"`
	Players       []Player `json:"players"`
	Results       []Result `json:"results"`
}

func ScrapeJunni(url string) (*JunniScrapingResult, error) {
	raw, err := scrapeJunniRaw(url)
	if err != nil {
		return nil, err
	}
	return ParseRawJunni(raw)
}

func ParseRawJunni(raw *JunniScrapingRawResult) (*JunniScrapingResult, error) {
	result := JunniScrapingResult{}

	result.URL = raw.URL
	result.RetrievalTime = raw.RetrievalTime
	result.Name = raw.Name
	result.HashAlgorithm = raw.HashAlgorithm
	result.Hash = raw.Hash
	result.Players = make([]Player, 0, len(raw.Table))
	names := make([]string, 0, len(raw.Table))
	for _, row := range raw.Table {
		if len(row) >= 2 {
			player := Player{ID: 0, Name: row[2][0]}
			result.Players = append(result.Players, player)
			names = append(names, player.Name)
		}
	}
	identifier := nameIdentifier{names: names, cache: make(map[string]int)}
	for i, row := range raw.Table {
		if len(row) >= 5 {
			for j := 5; j < len(row); j++ {
				if row[j][0] == "－" {
					continue
				}
				opponent := identifier.getIndex(row[j][1])
				if opponent == -1 {
					return nil, fmt.Errorf("opponent not found: (%d, %d): %s", i, j, row[j][1])
				}
				if i > opponent {
					continue
				}
				winner := 0
				playFirst := 0
				if row[j][0] == "○" {
					winner = 1
				} else if row[j][0] == "●" {
					winner = 2
				} else if row[j][0] == "先" {
					playFirst = 1
				} else if row[j][0] == "" {
					playFirst = 2
				} else {
					return nil, fmt.Errorf("invalid result: (%d, %d): %s", i, j, row[j][0])
				}
				result.Results = append(result.Results, Result{
					Index1:    i,
					Index2:    opponent,
					Winner:    winner,
					PlayFirst: playFirst,
					Round:     j - 4,
				})
			}
		}
	}
	sort.Sort(results(result.Results))

	return &result, nil
}
