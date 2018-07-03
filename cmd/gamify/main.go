package main

import (
	"fmt"
	"log"

	"github.com/softsrv/gamify/internal/pkg/router"
)

func main() {
	var myRouter router.Service
	fmt.Println("Starting listener on :8080")
	log.Fatal(myRouter.Start(":8080"))
}
