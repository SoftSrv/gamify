package main

import (
	"fmt"
	"log"

	"github.com/softsrv/gamify/internal/pkg/router"
)

func main() {
	fmt.Println("hello world")

	var myRouter router.Service

	log.Fatal(myRouter.Start(":8080"))
}
