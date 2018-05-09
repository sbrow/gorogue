package main

import (
	termbox "github.com/nsf/termbox-go"
	rogue "github.com/sbrow/gorogue"
)

func main() {
	host, port := "localhost", ":6061"
	_ = rogue.NewServer(host, port, rogue.NewMap(100, 100, "Map"))
	client := rogue.Connect(host, port)

	p := rogue.NewPlayer("Player", &rogue.Point{3, 3},
		termbox.Cell{'1', termbox.ColorWhite, termbox.ColorDefault}, 1)

	err := client.Spawn(p)
	if err != nil {
		panic(err)
	}
	rogue.CurrUI.Run()
}
