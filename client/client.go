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

func Move(a Actor, x, y int) {
	var args *MoveArgs
	var reply *bool
	args = &MoveArgs{Actors([]Actor{a}), []Point{Point{x, y}}}
	err := std.client.Call("Server.Move", args, &reply)
	if err != nil {
		panic(err)
	}
}

func Spawn(a ...Actor) error {
	var reply *SpawnReply
	err := std.client.Call("Server.Spawn", Actors(a), &reply)
	ui = Fullscreen(reply.Map).UI
	std.Squad = reply.Actors
	return err
}
