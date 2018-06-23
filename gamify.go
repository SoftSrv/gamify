package main

import (
	"fmt"
	"log"
	"net/http"

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
	var x mockapi.Service
	pl, err := x.Players("fakeid")
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "failed to fetch players")
		return
	}
	fmt.Println("here we go")
	fmt.Fprintf(w, "%+v\n", pl)

}

// Friends accepts one steamID in the query and
// returns all friends of that steamID
func Friends(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//queryValues := r.URL.Query()
}

// Games accepts one steamID in the query, and returns
// all games owned by that steam user
func Games(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//queryValues := r.URL.Query()
	// to construct url for game icon, use:
	// http://media.steampowered.com/steamcommunity/public/images/apps/{appid}/{hash}.jpg
}

func main() {
	fmt.Println("hello world")

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/players", Players)
	router.GET("/friends", Friends)
	router.GET("/games", Games)

	log.Fatal(http.ListenAndServe(":8080", router))
}
