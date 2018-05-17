// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package single_player

import (
	"errors"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/assets"
	"github.com/sbrow/gorogue/example"
	"log"
)

type Client struct {
	ui    engine.UI
	world *example.World
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
	a := engine.NewAction("Spawn", c.Addr(), example.NewPlayer("Player"))
	if err := c.HandleAction(a); err != nil {
		log.Println(err)
	}
	c.ui = assets.Fullscreen(c, c.world.Maps()[0])
	return nil
}

func (c *Client) Maps() []engine.Map {
	var maps []engine.Map
	for i, m := range c.world.Maps() {
		maps[i] = m
	}
	return maps
}

func (c *Client) Run() {
	c.ui.Run()
}
func (c *Client) Player() engine.Actor {
	return c.world.Players()[c.Addr()]
}

func (c *Client) SetWorld(w engine.World) {
	c.world = w.(*example.World)
}
