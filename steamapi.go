package steamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	apiKey string
}

// NewService returns a service object configured with the provided Steam web API Key
func NewService(apiKey string) *Service {
	return &Service{
		client: &http.Client{},
		apiKey: apiKey,
	}
}

// Players accepts one or more steamIDs and returns a PlayersResult
func (s *Service) Players(steamIDs []string) (*PlayersResult, error) {
	fmt.Printf("starting steam PLAYERS call for steamID list: %s\n", steamIDs)
	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/"+userService+"/GetPlayerSummaries/v0002?"+url.Values{
			"steamids": {strings.Join(steamIDs, ",")},
			"key":      {s.apiKey},
		}.Encode(),
		nil,
	)
	if err != nil {
		fmt.Println("got an error forming the request")
		return nil, err
	}
	fmt.Println("successfully formed the http request")
	req.Header.Add("Content-Type", "application/json")
	fmt.Println("just added the content type header")
	res, err := s.client.Do(req)
	if res != nil {
		fmt.Println("got a response, so defering the body close")
		defer res.Body.Close()
	}

	if err != nil {
		fmt.Println("got an error on .Do of http request")
		return nil, err
	}
	fmt.Println("no error received in http request")
	parsedPlayers := PlayersResult{}
	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(data))
	if err := json.NewDecoder(res.Body).Decode(&parsedPlayers); err != nil {
		fmt.Println("got an error decoding the http response body")
		fmt.Printf("the err is: %s\n", err)
		return nil, err
	}
	fmt.Println("no errors received in PLAYERS steam call")
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
