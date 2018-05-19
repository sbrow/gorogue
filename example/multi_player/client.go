// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package example

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
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

type RemoteClient struct {
	addr   string
	client *rpc.Client
	squad  Actors
	maps   map[string]*Map
}

func (r *RemoteClient) Addr() string {
	return r.addr
}

func (r *RemoteClient) Disconnect() {
	var reply *string
	addr := r.Addr()
	if err := r.client.Call("Server.Disconnect", &addr, &reply); err != nil {
		panic(err)
	}
	if err := r.client.Close(); err != nil {
		panic(err)
	}
	log.Println("Disconnected from server.")
}

func (r *RemoteClient) HandleAction(a *engine.Action) error {
	switch a.Name() {
	case "Quit":
		return errors.New("Leaving...")
	case "Move":
		return r.Move(a)
	default:
		return nil
	}
}

func (r *RemoteClient) Init() engine.UI {
	r.Spawn(NewPlayer("Player"))
	r.Ping()
	// TODO: (10) active squad member
	return assets.Fullscreen(r, &r.Squad()[0].Pos().Map)
}

func (r *RemoteClient) Maps() map[string]engine.Map {
	m := map[string]engine.Map{}
	for k, v := range r.maps {
		m[k] = v
	}
	return m
}

// Move requests that the server move actor a in direction dir.
func (r *RemoteClient) Move(a *engine.Action) error {
	var ma engine.Move
	// TODO: Make converter from Action to *Move
	var p engine.Pos
	if a.Name() != "Move" {
		//TODO: ErrorWrongAction or something.
		return nil
	}
	if a.Caller() == "Client" {
		caller := r.Squad()[0] // TODO: (8) Implement active squad member.
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
	err := r.client.Call("Server.Move", ma, &reply)
	CheckErrors(err)
	return nil
}

// Ping asks the server for all information relevant to the client.
func (r *RemoteClient) Ping() {
	var reply *Pong
	err := r.client.Call("Server.Ping", r.Addr(), &reply)
	CheckErrors(err)
	r.maps = reply.Maps
	r.squad = reply.Squad
}

func (r *RemoteClient) SetAddr(addr string) {
	r.addr = addr
}

func (r *RemoteClient) SetRPC(conn *rpc.Client) {
	r.client = conn
}

// Spawn requests that the server spawn actors.
// The server determines where to spawn them and returns the map
// where they spawned.
func (r *RemoteClient) Spawn(a ...engine.Actor) {
	args := &Spawn{
		Caller: r.Addr(),
		Actors: a,
	}
	var reply *bool
	err := r.client.Call("Server.Spawn", args, &reply)
	CheckErrors(err)
}

func (r *RemoteClient) Squad() []engine.Actor {
	return r.squad
}
