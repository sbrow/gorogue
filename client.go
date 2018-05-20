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

func (c *ExampleClient) Run() {
	c.UI.Run()
}

func (c *ExampleClient) Player() Actor {
	return c.World.Players()[c.Addr()]
}
