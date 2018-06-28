package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/softsrv/gamify/steam"
)

//Index displays a basic welcome page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Players returns a list of steam players to the caller
func Players(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	var x mockapi.Service
	pl, err := x.Players("fake,id,string")
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch players")
		return
	}
	fmt.Println("here we go")
	json.NewEncoder(w).Encode(pl)

}

// Player returns a single steam player to the caller based on steamID
func Player(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	var x mockapi.Service
	pl, err := x.Player(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch players")
		return
	}
	fmt.Println("here we go")
	json.NewEncoder(w).Encode(pl)

}

// Friends accepts one or more steamIDs in the query and
// returns all friends of that steamID
func Friends(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//queryValues := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	var x mockapi.Service
	fr, err := x.Friends(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch friends")
		return
	}
	fmt.Println("here we go")
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
	// http://media.steampowered.com/steamcommunity/public/images/apps/{appid}/{hash}.jpg
	w.Header().Set("Content-Type", "application/json")
	var x mockapi.Service
	ga, err := x.Games(params.ByName("id"))
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch games")
		return
	}
	fmt.Println("here we go")
	json.NewEncoder(w).Encode(ga)
}

func main() {
	fmt.Println("hello world")

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/players", Players)
	router.GET("/players/:id", Player)
	router.GET("/players/:id/friends", Friends)
	router.GET("/players/:id/games", Games)

	log.Fatal(http.ListenAndServe(":8080", router))
}
