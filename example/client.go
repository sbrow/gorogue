// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package example

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"net/rpc"
	"os"
)

func CheckError(e error) {
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

type Client struct {
	addr   string
	client *rpc.Client
	squad  Actors
	maps   map[string]*Map
}

func (c *Client) Addr() string {
	return c.addr
}

func (c *Client) HandleAction(a engine.Action) {
	switch a.Name() {
	case "Quit":
		termbox.Close()
		os.Exit(0)
	case "Move":
		c.Move(a)
	}
}

func (c *Client) Init() *engine.UI {
	c.Spawn(NewPlayer("Player"))
	c.Ping()
	// TODO: (10) active squad member
	return engine.Fullscreen(&c.Squad()[0].Pos().Map).UI
}

func (c *Client) Maps() map[string]engine.Map {
	m := map[string]engine.Map{}
	for k, v := range c.maps {
		m[k] = v
	}
	return m
}

// Move requests that the server move actor a in direction dir.
func (c *Client) Move(a engine.Action) {
	var ma MoveAction
	// TODO: Make converter from Action to *MoveAction
	var p engine.Pos
	if a.Name() != "Move" {
		return
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

	var reply *engine.ActionResponse
	err := c.client.Call("Server.Move", ma, &reply)
	CheckError(err)
}

// Ping asks the server for all information relevant to the client.
func (c *Client) Ping() {
	var reply *Pong
	err := c.client.Call("Server.Ping", c.Addr(), &reply)
	CheckError(err)
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
	args := &SpawnAction{
		Caller: c.Addr(),
		Actors: a,
	}
	var reply *bool
	// Unsafe
	// _ = c.client.Call("Server.Spawn", args, &reply)

	// Safe
	err := c.client.Call("Server.Spawn", args, &reply)
	CheckError(err)
}

func (c *Client) Squad() []engine.Actor {
	return c.squad
}
