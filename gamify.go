package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Index displays a basic welcome message
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
}

func User(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func Games(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
}

func main() {
	fmt.Println("hello world")

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users", Users)
	router.GET("/users/:steamID", User)
	router.GET("/games", Games)

	log.Fatal(http.ListenAndServe(":8080", router))
}
