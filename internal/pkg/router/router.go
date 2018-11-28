package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/softsrv/gamify/pkg/steam"
)

// fill in with /{appId}/{hash}
const gameIconURL = "http://media.steampowered.com/steamcommunity/public/images/apps/%d/%s.jpg"

// Service provides access to the route handler
type Service struct {
	router *httprouter.Router
}

//Start defines the router, sets up the routes, and starts the listener
func (s *Service) Start(port string) error {
	s.router = httprouter.New()
	s.setRoutes()

	return http.ListenAndServe(port, s.router)
}

func (s *Service) setRoutes() {

	s.router.GET("/", Index)
	s.router.GET("/players", Players)
	s.router.GET("/players/:id", Player)
	s.router.GET("/players/:id/friends", Friends)
	s.router.GET("/players/:id/games", Games)

}

//Index displays a basic welcome page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Players returns a list of steam players to the caller
func Players(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	x := steamapi.NewService(os.Getenv("STEAM_WEBAPI_KEY"))
	pl, err := x.Players("xxx,yyy,zzz") // will come from query params in the future
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch players")
		return
	}
	json.NewEncoder(w).Encode(pl)

}

// Player returns a single steam player to the caller based on steamID
func Player(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	var x steamapi.Service
	pl, err := x.Player(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch players")
		return
	}
	json.NewEncoder(w).Encode(pl)

}

// Friends accepts one or more steamIDs in the query and
// returns all friends of that steamID
func Friends(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	var x steamapi.Service
	fr, err := x.Friends(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch friends")
		return
	}
	var fID []string
	for _, item := range fr.Response.Friends {
		fID = append(fID, item.SteamID)
	}
	finalResult, err := x.Players(strings.Join(fID[:], ","))
	json.NewEncoder(w).Encode(finalResult)
}

// Games accepts one steamID in the query, and returns
// all games owned by that steam user
func Games(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//queryValues := r.URL.Query()
	// to construct url for game icon, use:
	//
	w.Header().Set("Content-Type", "application/json")
	var x steamapi.Service
	ga, err := x.Games(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch games")
		return
	}
	for i := 0; i < len(ga.Response.Games); i++ {
		game := &ga.Response.Games[i]
		game.ImgIconURL = fmt.Sprintf(gameIconURL, game.AppID, game.ImgIconURL)
		game.ImgLogoURL = fmt.Sprintf(gameIconURL, game.AppID, game.ImgLogoURL)
	}

	json.NewEncoder(w).Encode(ga)
}
