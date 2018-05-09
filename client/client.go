// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package client

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// std is the active client. Each process can have only one active client.
var std *client

// ui is the active user interface. Each process can have only one active UI.
var ui *UI

type client struct {
	client *rpc.Client
	Squad  Actors
}

// Connect initializes a connection to a server. It must be called before all other
// functions.
func Connect(host, port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		panic(err)
	}
	std = &client{
		client: jsonrpc.NewClient(conn),
		Squad:  Actors(Actors{}),
	}
}

// GetMap asks the server for the map our Squad is on.
//
// TODO: Tick shenanigans.
func GetMap(name string) Map {
	var reply *Map
	err := std.client.Call("Server.Map", name, &reply)
	if err != nil {
		panic(err)
	}
	return *reply
}

// Move requests that the server move actor a in direction dir.
func Move(a Actor, dir Direction) {
	var args *MoveArgs
	var reply *bool
	args = &MoveArgs{Actors([]Actor{a}), []Point{*a.Pos()}}

	if dir&North == North {
		args.Points[0].Y--
	} else {
		if dir&South == South {
			args.Points[0].Y++
		}
	}
	if dir&East == East {
		args.Points[0].X++
	} else {
		if dir&West == West {
			args.Points[0].X--
		}
	}
	err := std.client.Call("Server.Move", args, &reply)
	if *reply {
		a.SetPos(args.Points[0])
	}
	if err != nil {
		panic(err)
	}
}

// Spawn requests that the server spawn actors.
// The server determines where to spawn them and returns the map
// where they spawned.
func Spawn(a ...Actor) Actors {
	var reply *SpawnReply
	err := std.client.Call("Server.Spawn", Actors(a), &reply)
	if err != nil {
		panic(err)
	}
	ui = Fullscreen(reply.Map).UI
	std.Squad = reply.Actors
	return std.Squad
}
