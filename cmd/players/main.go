package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

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

func main() {
	var currentPlayers players.PlayerList
	currentPlayersBytes, err := os.ReadFile(PlayersFilename)
	needsUpdate := true
	if err == nil {
		json.Unmarshal(currentPlayersBytes, &currentPlayers)
		lastUpdateTime, err := time.Parse(time.RFC3339, currentPlayers.RetrievalTime)
		if err != nil {
			panic(err)
		}
		if lastUpdateTime.Add(24 * time.Hour).After(time.Now()) {
			needsUpdate = false
		}
	}
	if needsUpdate {
		updatePlayers(&currentPlayers)
		playersJSON, err := json.MarshalIndent(currentPlayers, "", "  ")
		if err != nil {
			panic(err)
		}
		os.WriteFile(PlayersFilename, playersJSON, 0o644)
	}
	// ensure the directory player-info/ exists
	err = os.Mkdir("player-info", 0o755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for _, player := range currentPlayers.Players {
		playerInfoFilename := fmt.Sprintf("player-info/%s-%d.json", player.Kind, player.ID)
		var currentPlayerInfo players.PlayerInfoRaw
		currentPlayerInfoBytes, err := os.ReadFile(playerInfoFilename)
		needsUpdate := true
		if err == nil {
			json.Unmarshal(currentPlayerInfoBytes, &currentPlayerInfo)
			lastUpdateTime, err := time.Parse(time.RFC3339, currentPlayerInfo.RetrievalTime)
			if err != nil {
				panic(err)
			}
			if lastUpdateTime.Add(24 * time.Hour).After(time.Now()) {
				needsUpdate = false
			}
		}
		if needsUpdate {
			url := fmt.Sprintf("https://www.shogi.or.jp/player/%s/%d.html", player.Kind, player.ID)
			playerInfo, err := players.ScrapePlayerInfoRaw(url, &player, &currentPlayers)
			if err != nil {
				panic(err)
			}
			currentPlayerInfo = *playerInfo

			playerInfoJSON, err := json.MarshalIndent(currentPlayerInfo, "", "  ")
			if err != nil {
				panic(err)
			}
			os.WriteFile(playerInfoFilename, playerInfoJSON, 0o644)
		}
	}
}
