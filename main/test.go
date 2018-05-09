package main

import (
	termbox "github.com/nsf/termbox-go"
	rogue "github.com/sbrow/gorogue"
)

func main() {
	p := rogue.NewPlayer("Player", &rogue.Point{3, 3},
		termbox.Cell{'@', termbox.ColorWhite, termbox.ColorDefault}, 1)
	a := rogue.Actors{p}
}
