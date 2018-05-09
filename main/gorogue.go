package main

import (
	termbox "github.com/nsf/termbox-go"
	rogue "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/client"
	server "github.com/sbrow/gorogue/server"
)

func main() {
	host, port := "localhost", ":6061"
	_ = server.Start(host, port, rogue.NewMap(100, 100, "Map"))
	client.Connect(host, port)

	p := rogue.NewPlayer("Player",
		termbox.Cell{'1', termbox.ColorWhite, termbox.ColorDefault}, 1)

	client.Spawn(p)
	client.Run()
}
