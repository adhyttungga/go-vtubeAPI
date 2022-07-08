package main

import (
	"github.com/adhyttungga/go-vtubeAPI/connection"
	"github.com/adhyttungga/go-vtubeAPI/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleRequest()
}