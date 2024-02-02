package main

import (
	"encoding/json"
	"os"

	"github.com/koba-e964/go-shogi-scraper/players"
)

func main() {
	playersList, err := players.ScrapePlayerList(players.DefaultURL)
	if err != nil {
		panic(err)
	}
	playersJSON, err := json.MarshalIndent(playersList, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("players.json", playersJSON, 0o644)
}
