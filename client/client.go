// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package client

import (
	"bytes"
	"fmt"
	. "github.com/sbrow/gorogue"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// std is the active client. Each process can have only one active client.
var std *client

// ui is the active user interface. Each process can have only one active UI.
var ui *UI

type client struct {
	Addr   string
	client *rpc.Client
	Squad  Actors
	Maps   map[string]*Map
}

func Check(e error) {
	switch e {
	case rpc.ErrShutdown:
		fmt.Println("The server shut down unexpectedly.")
		os.Exit(1)
	case nil:
		return
	default:
		panic(e)
	}
}

// Connect initializes a connection to a server. It must be called before all other
// functions.
func Connect(host, port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var addr []byte = make([]byte, 24)
	if _, err := conn.Read(addr); err != nil {
		panic(err)
	}
	std = &client{
		Addr:   string(bytes.Trim(addr, "\x00")),
		client: jsonrpc.NewClient(conn),
		Squad:  []Actor{},
		Maps:   map[string]*Map{},
	}
	Spawn(NewPlayer("Player", 1))
	Ping()
	// TODO: (10) active squad member
	fmt.Printf("%+v\n", std)
	ui = Fullscreen(&std.Squad[0].Pos().Map).UI
	Run()
}

// Ping asks the server for relevant information.
func Ping() {
	var reply *Pong
	err := std.client.Call("Server.Ping", std.Addr, &reply)
	fmt.Printf("%+v\n", reply.Maps["Map"])
	Check(err)
	std.Maps = reply.Maps
	std.Squad = reply.Squad
}

// Move requests that the server move actor a in direction dir.
func Move(a *Action) {
	var ma MoveAction
	var p Pos
	if a.Name != "Move" {
		return
	}
	if a.Caller == "Client" {
		caller := std.Squad[0] // TODO: (8) Implement active squad member.
		ma.Caller = caller.Name()
		p = *caller.Pos()
	}
	switch a.Args[0].(type) {
	case Direction:
		dir := a.Args[0].(Direction)
		if dir&North == North {
			p.Y--
		} else {
			if dir&South == South {
				p.Y++
			}
		}
		if dir&East == East {
			p.X++
		} else {
			if dir&West == West {
				p.X--
			}
		}
	case Pos:
		p = a.Args[0].(Pos)
	default:
		panic("Passed wrong args to Client.Move()")
	}
	ma.Pos = p

	var reply *ActionResponse
	err := std.client.Call("Server.Move", ma, &reply)
	Check(err)
}

// Spawn requests that the server spawn actors.
// The server determines where to spawn them and returns the map
// where they spawned.
func Spawn(a ...Actor) {
	args := &SpawnAction{
		Caller: std.Addr,
		Actors: a,
	}
	var reply *bool
	err := std.client.Call("Server.Spawn", args, &reply)
	Check(err)
}
