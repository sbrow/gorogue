package client

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var std *Client
var ui *UI

type Client struct {
	client *rpc.Client
	Squad  Actors
}

func Connect(host, port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		panic(err)
	}
	std = &Client{
		client: jsonrpc.NewClient(conn),
		Squad:  Actors(Actors{}),
	}
}

func GetMap(name string) Map {
	var reply *Map
	err := std.client.Call("Server.Map", name, &reply)
	if err != nil {
		panic(err)
	}
	return *reply
}

func Move(a Actor, dir Direction) {
	var args *MoveArgs
	var reply *bool
	args = &MoveArgs{Actors([]Actor{a}), []Point{*a.Pos()}}

	switch dir {
	case North:
		args.Points[0].Y--
	case East:
		args.Points[0].X++
	case West:
		args.Points[0].X--
	case South:
		args.Points[0].Y++
	case NorthEast:
		args.Points[0].X++
		args.Points[0].Y--
	}
	err := std.client.Call("Server.Move", args, &reply)
	if *reply {
		a.SetPos(args.Points[0])
	}
	if err != nil {
		panic(err)
	}
}

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
