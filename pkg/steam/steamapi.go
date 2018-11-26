package steamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const dataLocation = "internal/pkg/mock/mockdata"
const baseURL = "https://api.steampowered.com"
const userService = "ISteamUser"
const playerService = "IPlayerService"

// Player contains details about the Steam User
type Player struct {
	SteamID      string `json:"steamid"`
	PersonaName  string `json:"personaname"`
	AvatarSmall  string `json:"avatar"`
	AvatarMedium string `json:"avatarmedium"`
	AvatarFull   string `json:"avatarfull"`
}

// PlayersList contains an array of Player objects.
type PlayersList struct {
	Players []Player `json:"players"`
}

// PlayersResult contains a "response" object with relevant data
type PlayersResult struct {
	Response PlayersList `json:"response"`
}

// Game contains details about a Steam game
type Game struct {
	AppID           int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconURL      string `json:"img_icon_url"`
	ImgLogoURL      string `json:"img_logo_url"`
}

// GamesList contains an array of Game objects
type GamesList struct {
	Games []Game `json:"games"`
}

// GamesResult contains a "response" object with relevant data
type GamesResult struct {
	Response GamesList `json:"response"`
}

// A Friend is a reference to a Player who is friends with a particular user
type Friend struct {
	SteamID     string `json:"steamid"`
	FriendSince int    `json:"friend_since"`
}

// FriendsList contains an array of Friend objects
type FriendsList struct {
	Friends []Friend `json:"friends"`
}

// FriendsResult contains a "response" object with relevant data
type FriendsResult struct {
	Response FriendsList `json:"response"`
}

// Service is the type that owns methods for fetching steam data
type Service struct {
	client *http.Client
}

// Players accepts one or more steamIDs and returns a PlayersResult
func (s *Service) Players(steamID string) (*PlayersResult, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/"+userService+"/"+"/GetPlayerSummaries/v0002"+url.Values{
			"steamids": {steamID},
			"key":      {os.Getenv("STEAM_WEBAPI_KEY")},
		}.Encode(),
		nil,
	)
	req.Header.Add("Content-Type", "application/json")
	rawPlayers, err := ioutil.ReadFile(dataLocation + "/players.json")
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
func (s *Service) Player(steamID string) (*Player, error) {
	rawPlayers, err := ioutil.ReadFile(dataLocation + "/players.json")
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
func (s *Service) Games(steamIDs string) (*GamesResult, error) {
	rawGames, err := ioutil.ReadFile(dataLocation + "/games.json")
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
func (s *Service) Friends(steamIDs string) (*FriendsResult, error) {
	rawFriends, err := ioutil.ReadFile(dataLocation + "/friends.json")
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
