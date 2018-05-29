package gorogue

import (
	termbox "github.com/nsf/termbox-go"
)

var stdConn Client

// Commonly used roguelike symbols.
var (
	DefaultPlayer = termbox.Cell{Ch: '@', Fg: termbox.ColorWhite, Bg: termbox.ColorBlack}
	StairsUp      = termbox.Cell{Ch: '>', Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}
	StairsDown    = termbox.Cell{Ch: '<', Fg: termbox.ColorDefault, Bg: termbox.ColorDefault}
)
