package gorogue

import (
	"errors"
)

type ExampleClient struct {
	UI    UI
	World *World
}

func (c *ExampleClient) Addr() string {
	return "[::1]"
}

func (c *ExampleClient) HandleAction(a *Action) error {
	var reply *string

	switch a.Name {
	case "Quit":
		return errors.New("Leaving...")
	default:
		if a.Caller == "Client" {
			a.Caller = c.Addr()
		}
		c.World.HandleAction(a, reply)
	}
	if reply != nil {
		return errors.New(*reply)
	}
	return nil
}

func (c *ExampleClient) Init() error {
	// Create a new Map.
	m := NewMap(5, 5, "Map_1")

	// Add it to our world.
	c.World = NewWorld(m)

	a := NewAction("Spawn", c.Addr(), NewPlayer("Player"))
	if err := c.HandleAction(a); err != nil {
		Log.Println(err)
	}
	// c.UI = ui.Standard(c, c.World.Maps()[0])
	return nil
}

func (c *ExampleClient) Run() {
	c.UI.Run()
}

func (c *ExampleClient) Player() Actor {
	return c.World.Players[c.Addr()]
}
