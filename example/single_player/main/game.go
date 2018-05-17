package main

import (
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/example"
	"github.com/sbrow/gorogue/example/single_player"
)

func main() {
	// Set up our log
	f, err := engine.SetLog("local_game")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create a new Map.
	m := example.NewMap(15, 24, "Map_1")

	// Add it to our world.
	w := example.NewWorld(m)

	// Run the Client
	engine.NewClient(&single_player.Client{}, w)
}
