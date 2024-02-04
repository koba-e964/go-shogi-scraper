package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	scraper "github.com/koba-e964/go-shogi-scraper"
	"github.com/koba-e964/go-shogi-scraper/players"
)

const PlayersFilename = "players.json"

func updatePlayers(currentPlayers *players.PlayerList) error {
	playersList, err := players.ScrapePlayerList(players.DefaultURL)
	if err != nil {
		return err
	}
	*currentPlayers = *playersList
	return nil
}

func retrieveWithCache[T scraper.RetrievedData](filename string, retrieve func() T) (T, error) {
	var currentData T
	currentDataBytes, err := os.ReadFile(filename)
	if err == nil {
		if err := json.Unmarshal(currentDataBytes, &currentData); err == nil {
			lastUpdateTime, err := time.Parse(time.RFC3339, currentData.GetRetrievalTime())
			if err == nil {
				if lastUpdateTime.Add(24 * time.Hour).After(time.Now()) {
					return currentData, nil
				}
			}
		}
	}
	// needs update
	currentData = retrieve()
	jsonContent, err := json.MarshalIndent(currentData, "", "  ")
	if err != nil {
		return currentData, err
	}
	os.WriteFile(filename, jsonContent, 0o644)
	return currentData, nil
}

func main() {
	currentPlayers, err := retrieveWithCache(PlayersFilename, func() *players.PlayerList {
		var tmp players.PlayerList
		updatePlayers(&tmp)
		return &tmp
	})
	if err != nil {
		panic(err)
	}
	// ensure the directory player-info/ exists
	err = os.Mkdir("player-info", 0o755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for _, player := range currentPlayers.Players {
		playerInfoFilename := fmt.Sprintf("player-info/%s-%d.json", player.Kind, player.ID)
		_, err := retrieveWithCache(playerInfoFilename, func() *players.PlayerInfoRaw {
			url := fmt.Sprintf("https://www.shogi.or.jp/player/%s/%d.html", player.Kind, player.ID)
			playerInfo, err := players.ScrapePlayerInfoRaw(url, &player, currentPlayers)
			if err != nil {
				panic(err)
			}
			return playerInfo
		})
		if err != nil {
			panic(err)
		}
	}
}
