package main

import (
	"log"

	"github.com/softsrv/gamify/internal/pkg/router"
)

func main() {
	var myRouter router.Service
	log.Fatal(myRouter.Start(":8080"))
}
