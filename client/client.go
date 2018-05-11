// Package client handles drawing the UI, interfacing with the player,
// and talking to the server.
package client

import (
	"encoding/json"
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
	Squad  []Player
}

// Connect initializes a connection to a server. It must be called before all other
// functions.
func Connect(host, port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	std = &client{
		client: jsonrpc.NewClient(conn),
		Squad:  []Player{},
	}
	m, p := Spawn(NewPlayer("Player", 1))
	for _, a := range p {
		std.Squad = append(std.Squad, a.(Player))
	}
	ui = Fullscreen(m).UI
	Run()
}

// GetMap asks the server for the map our Squad is on.
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
	args := &Action{
		Caller: a.Name(),
	}
	p := a.Pos()

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
	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	args.Args = [][]byte{data}

	var reply *ActionResponse
	err = std.client.Call("Server.Move", args, &reply)
	if err != nil {
		panic(err)
	}
}

// Spawn requests that the server spawn actors.
// The server determines where to spawn them and returns the map
// where they spawned.
func Spawn(a ...Actor) (Map *string, Spawned Actors) {
	var reply *SpawnReply
	data, _ := json.Marshal(Actors(a))
	fmt.Println(string(data))
	err := std.client.Call("Server.Spawn", Actors(a), &reply)
	if err != nil {
		panic(err)
	}
	return reply.Map, reply.Actors
}
