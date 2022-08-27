package main

import (
	"GinBaseProject/api"
	"log"
)

func main() {
	var server = api.LoadUrls()

	log.Fatal(server.Run(":8080"))

}
