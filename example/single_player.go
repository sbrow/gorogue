package main

import (
	"github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/assets"
	"github.com/sbrow/gorogue/ui"
)

func main() {
	// Set up our log
	f, err := gorogue.SetLog("local_game.log", true)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c := &assets.ExampleClient{}
	gorogue.NewClient(c)
	assets.StandardUI(c.Map())

	// Run the Client
	ui.Run()
}
