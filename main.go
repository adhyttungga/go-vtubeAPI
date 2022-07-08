package main

import (
	"fmt"

	"github.com/adhyttungga/go-vtubeAPI/connection"
	"github.com/adhyttungga/go-vtubeAPI/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleRequest()

	fmt.Println("Server Up and Running on Port 11000")
}