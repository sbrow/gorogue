package main

import (
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
	"log"
)

type Client struct {
	engine.ExampleClient
}

func (c *Client) Init() error {
	// Create a new Map.
	m := engine.NewMap(5, 5, "Map_1")

	// Add it to our world.
	c.World = engine.NewWorld(m)

	a := engine.NewAction("Spawn", c.Addr(), engine.NewPlayer("Player"))
	if err := c.HandleAction(a); err != nil {
		log.Println(err)
	}
	c.UI = ui.Fullscreen(c, c.World.Maps()[0])
	return nil
}

func main() {
	// Set up our log
	f, err := engine.SetLog("local_game")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Run the Client
	engine.NewClient(&Client{})
}
