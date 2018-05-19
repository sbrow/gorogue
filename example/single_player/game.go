package main

import (
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/example"
)

func main() {
	// Set up our log
	f, err := engine.SetLog("local_game")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Run the Client
	engine.NewClient(&example.Client{})
}
