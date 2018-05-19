// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package example

import (
	"errors"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
	"log"
)

type Client struct {
	ui    engine.UI
	world *engine.World
}

func (c *Client) Addr() string {
	return "[::1]"
}

func (c *Client) HandleAction(a *engine.Action) error {
	var reply *string
	switch a.Name {
	case "Quit":
		return errors.New("Leaving...")
	default:
		if a.Caller == "Client" {
			a.Caller = c.Addr()
		}
		c.world.HandleAction(a, reply)
	}
	if reply != nil {
		return errors.New(*reply)
	}
	return nil
}

func (c *Client) Init() error {
	// Create a new Map.
	m := engine.NewMap(100, 100, "Map_1")

	// Add it to our world.
	c.world = engine.NewWorld(m)

	a := engine.NewAction("Spawn", c.Addr(), engine.NewPlayer("Player"))
	if err := c.HandleAction(a); err != nil {
		log.Println(err)
	}
	c.ui = ui.Fullscreen(c, c.world.Maps()[0])
	return nil
}

func (c *Client) Run() {
	c.ui.Run()
}
func (c *Client) Player() engine.Actor {
	return c.world.Players()[c.Addr()]
}
