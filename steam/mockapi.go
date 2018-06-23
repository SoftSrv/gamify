package mockapi

import (
	"encoding/json"
	"io/ioutil"
)

type Service struct {
}

// Player contains details about the Steam User
type Player struct {
	SteamID      string `json:"steamid"`
	PersonaName  string `json:"personaname"`
	AvatarSmall  string `json:"avatar"`
	AvatarMedium string `json:"avatarmedium"`
	AvatarFull   string `json:"avatarfull"`
}

type PlayerList struct {
	Players []Player `json:"players"`
}

type PlayersResult struct {
	Response PlayerList `json:"response"`
}

// Players accepts one or more steamIDs and returns a PlayersResult
func (m *Service) Players(steamIDs string) (*PlayersResult, error) {
	rawFriends, err := ioutil.ReadFile("steam/mockdata/players.json")
	if err != nil {
		return nil, err
	}
	var parsedFriends PlayersResult
	parseErr := json.Unmarshal(rawFriends, &parsedFriends)
	if parseErr != nil {
		return nil, parseErr
	}
	return &parsedFriends, nil
}
