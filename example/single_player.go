package main

import (
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
)

func main() {
	// Set up our log
	f, err := SetLog("local_game")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c := &ExampleClient{}
	c.Init()
	ui.Standard(c, c.World.Maps[0])

	// Run the Client
	ui.Run()
}
