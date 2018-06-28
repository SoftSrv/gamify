package mockapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

// Game contains details about a Steam game
type Game struct {
	AppID           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconURL      string `json:"img_icon_url"`
	ImgLogoURL      string `json:"img_logo_url"`
}

type GameList struct {
	Games []Game `json:"games"`
}

type GamesResult struct {
	Response GameList `json:"response"`
}

// A Friend is a reference to a Player
type Friend struct {
	SteamID     string `json:"steamid"`
	FriendSince string `json:"friend_since"`
}

type FriendList struct {
	Friends []Friend `json:"friends"`
}

type FriendsResult struct {
	Response FriendList `json:"friendslist"`
}

type Service struct {
}

// Players accepts one or more steamIDs and returns a PlayersResult
func (m *Service) Players(steamID string) (*PlayersResult, error) {
	rawPlayers, err := ioutil.ReadFile("steam/mockdata/players.json")
	if err != nil {
		return nil, err
	}
	var parsedPlayers PlayersResult
	parseErr := json.Unmarshal(rawPlayers, &parsedPlayers)
	if parseErr != nil {
		return nil, parseErr
	}
	return &parsedPlayers, nil
}

// Player accepts one steamID and returns that player's object
func (m *Service) Player(steamID string) (*Player, error) {
	rawPlayers, err := ioutil.ReadFile("steam/mockdata/players.json")
	if err != nil {
		return nil, err
	}
	var parsedPlayers PlayersResult
	parseErr := json.Unmarshal(rawPlayers, &parsedPlayers)
	if parseErr != nil {
		return nil, parseErr
	}

	for _, item := range parsedPlayers.Response.Players {
		if steamID == item.SteamID {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("couldnt find player for ID: %s", steamID)
}

// Games accepts one or more steamIDs and returns a GamesResult
func (m *Service) Games(steamIDs string) (*GamesResult, error) {
	rawGames, err := ioutil.ReadFile("steam/mockdata/games.json")
	if err != nil {
		return nil, err
	}
	var parsedGames GamesResult
	parseErr := json.Unmarshal(rawGames, &parsedGames)
	if parseErr != nil {
		return nil, parseErr
	}
	return &parsedGames, nil
}

// Friends accepts a steamID and returns all friends for that ID
func (m *Service) Friends(steamIDs string) (*FriendsResult, error) {
	rawFriends, err := ioutil.ReadFile("steam/mockdata/friends.json")
	if err != nil {
		return nil, err
	}
	var parsedFriends FriendsResult
	parseErr := json.Unmarshal(rawFriends, &parsedFriends)
	if parseErr != nil {
		return nil, parseErr
	}
	return &parsedFriends, nil
}
