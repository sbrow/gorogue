package gorogue

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var CurrClient *Client
var CurrUI *UI

type Client struct {
	client *rpc.Client
	Squad  Actors
}

func Connect(host, port string) *Client {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		panic(err)
	}
	CurrClient = &Client{
		client: jsonrpc.NewClient(conn),
		Squad:  Actors(Actors{}),
	}
	return CurrClient
}

func (c *Client) Map(name string) Map {
	var reply *Map
	err := c.client.Call("Server.Map", name, &reply)
	if err != nil {
		panic(err)
	}
	return *reply
}

func (c *Client) Move(a Actor, x, y int) {
	var args *MoveArgs
	var reply *bool
	args = &MoveArgs{Actors([]Actor{a}), []Point{Point{x, y}}}
	err := c.client.Call("Server.Move", args, &reply)
	if err != nil {
		panic(err)
	}
}

func (c *Client) Spawn(a ...Actor) error {
	var reply *SpawnReply
	err := c.client.Call("Server.Spawn", Actors(a), &reply)
	CurrUI = Fullscreen(&reply.Map).UI
	c.Squad = reply.Actors
	return err
}
