package assets

import (
	"errors"
	. "github.com/sbrow/gorogue"
)

type ExampleClient struct {
	UI    UI
	World *ExampleWorld
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
	// Create a new world.
	c.World = NewWorld()

	// Add a map to our world.
	c.World.NewMap(5, 5)

	a := NewAction("Spawn", c.Addr(), NewPlayer("Player"))
	if err := c.HandleAction(a); err != nil {
		Log.Println(err)
	}
	return nil
}

func (c *ExampleClient) Player() Actor {
	return c.World.Players[c.Addr()]
}
