package assets

import (
	"errors"
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/ui"
	"log"
	"net/rpc"
)

type ExampleRemoteClient struct {
	mp     *[][]Tile
	addr   string
	client *rpc.Client
}

func (e *ExampleRemoteClient) Addr() string {
	return e.addr
}

func (e *ExampleRemoteClient) Disconnect() {
	var reply *string
	addr := e.addr
	if err := e.client.Call("Server.Disconnect", &addr, &reply); err != nil {
		log.Println("error (d/c):", err)
		return
	}
	if err := e.client.Close(); err != nil {
		log.Println("error (d/c):", err)
		return
	}
	log.Println("Disconnected from server.")
}

func (e *ExampleRemoteClient) HandleAction(a *Action) error {
	var reply *string

	switch a.Name {
	case "Quit":
		return errors.New("Leaving...")
	default:
		if a.Caller == "Client" {
			a.Caller = e.Addr()
		}
		if err := e.client.Call("Server.HandleAction", &a, &reply); err != nil {
			return err
		}
	}
	if reply != nil && *reply != "" {
		return errors.New(*reply)
	}
	e.Ping()
	return nil
}

func (e *ExampleRemoteClient) Init() error {
	e.mp = &[][]Tile{}
	e.Ping()
	return nil
}

func (e *ExampleRemoteClient) Map() *[][]Tile {
	return e.mp
}

func (e *ExampleRemoteClient) Run() {
	ui.Standard(e.Map())
	ui.Run()
}

func (e *ExampleRemoteClient) Ping() error {
	args := 0
	if err := e.client.Call("Server.Map", &args, e.mp); err != nil {
		return err
	}
	return nil
}

func (e *ExampleRemoteClient) Player() Actor {
	return nil
}

func (e *ExampleRemoteClient) SetRPC(conn *rpc.Client) {
	e.client = conn
}

func (e *ExampleRemoteClient) SetAddr(addr string) {
	e.addr = addr
}
