package assets

import (
	"errors"
	. "github.com/sbrow/gorogue"
)

type ExampleClient struct {
	World *ExampleWorld
	mp    *[][]Tile
}

func (e *ExampleClient) Addr() string {
	return "[::1]"
}

func (e *ExampleClient) HandleAction(a *Action) error {
	var reply *string

	switch a.Name {
	case "Quit":
		return errors.New("Leaving...")
	default:
		if a.Caller == "Client" {
			a.Caller = e.Addr()
		}
		e.World.HandleAction(a, reply)
	}
	if reply != nil {
		return errors.New(*reply)
	}
	e.Ping()
	return nil
}

func (e *ExampleClient) Init() error {
	// Create a new world.
	e.World = NewWorld()

	// Add a map to our world.
	e.World.NewMap(5, 5)
	e.mp = &[][]Tile{}

	a := NewAction("Spawn", e.Addr(), NewPlayer("Player"))
	*e.mp = e.World.Maps()[0].AllTiles()
	if err := e.HandleAction(a); err != nil {
		Log.Println(err)
	}
	return nil
}

func (e *ExampleClient) Map() *[][]Tile {
	return e.mp
}

func (e *ExampleClient) Ping() error {
	*e.mp = e.World.Maps()[0].AllTiles()
	return nil
}

func (e *ExampleClient) Player() Actor {
	return e.World.Players()[e.Addr()]
}
