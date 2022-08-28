package main

import (
	"GinBaseProject/api"
	"log"
)

func main() {
	var server = api.LoadUrls()

	log.Fatal(server.Run("127.0.0.1:8080"))

}

type CreateTaskInputValidator struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}
