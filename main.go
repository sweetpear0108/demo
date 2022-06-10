package main

import (
	"awesomeProject/route"
	"log"
)

func main() {
	router := route.Init()
	log.Fatal(router.Run("localhost:8080"))
}
