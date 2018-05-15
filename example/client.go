// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package example

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/action"
	"github.com/sbrow/gorogue/assets"
	"log"
	"net/rpc"
	"os"
)

func CheckErrors(errs ...error) {
	for _, e := range errs {
		switch e {
		case rpc.ErrShutdown:
			termbox.Close()
			fmt.Println("The server shut down unexpectedly.")
			os.Exit(1)
		case nil:
			return
		default:
			panic(e)
		}
	}
}

type Client struct {
	addr   string
	client *rpc.Client
	squad  Actors
	maps   map[string]*Map
}

func (c *Client) Addr() string {
	return c.addr
}

func (c *Client) Disconnect() {
	var reply *string
	addr := c.Addr()
	if err := c.client.Call("Server.Disconnect", &addr, &reply); err != nil {
		panic(err)
	}
	if err := c.client.Close(); err != nil {
		panic(err)
	}
	log.Println("Disconnected from server.")
}

func (c *Client) HandleAction(a engine.Action) error {
	switch a.Name() {
	case "Quit":
		return errors.New("Leaving...")
	case "Move":
		return c.Move(a)
	default:
		return nil
	}
	return nil
}

func (c *Client) Init() engine.UI {
	c.Spawn(NewPlayer("Player"))
	c.Ping()
	// TODO: (10) active squad member
	return assets.Fullscreen(c, &c.Squad()[0].Pos().Map)
}

func (c *Client) Maps() map[string]engine.Map {
	m := map[string]engine.Map{}
	for k, v := range c.maps {
		m[k] = v
	}
	return m
}

// Move requests that the server move actor a in direction dir.
func (c *Client) Move(a engine.Action) error {
	var ma action.Move
	// TODO: Make converter from Action to *Move
	var p engine.Pos
	if a.Name() != "Move" {
		//TODO: ErrorWrongAction or something.
		return nil
	}
	if a.Caller() == "Client" {
		caller := c.Squad()[0] // TODO: (8) Implement active squad member.
		ma.Caller = caller.Name()
		p = *caller.Pos()
	}
	switch a.Args()[0].(type) {
	case engine.Direction:
		dir := a.Args()[0].(engine.Direction)
		if dir&engine.North == engine.North {
			p.Y--
		} else {
			if dir&engine.South == engine.South {
				p.Y++
			}
		}
		if dir&engine.East == engine.East {
			p.X++
		} else {
			if dir&engine.West == engine.West {
				p.X--
			}
		}
	case engine.Pos:
		p = a.Args()[0].(engine.Pos)
	default:
		panic("Passed wrong args to Client.Move()")
	}
	ma.Pos = p

	var reply *string
	err := c.client.Call("Server.Move", ma, &reply)
	CheckErrors(err)
	return nil
}

// Ping asks the server for all information relevant to the client.
func (c *Client) Ping() {
	var reply *Pong
	err := c.client.Call("Server.Ping", c.Addr(), &reply)
	CheckErrors(err)
	c.maps = reply.Maps
	c.squad = reply.Squad
}

func (c *Client) SetAddr(addr string) {
	c.addr = addr
}

func (c *Client) SetRPC(conn *rpc.Client) {
	c.client = conn
}

// Spawn requests that the server spawn actors.
// The server determines where to spawn them and returns the map
// where they spawned.
func (c *Client) Spawn(a ...engine.Actor) {
	args := &Spawn{
		Caller: c.Addr(),
		Actors: a,
	}
	var reply *bool
	err := c.client.Call("Server.Spawn", args, &reply)
	CheckErrors(err)
}

func (c *Client) Squad() []engine.Actor {
	return c.squad
}
