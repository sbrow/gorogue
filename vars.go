package gorogue

import (
	termbox "github.com/nsf/termbox-go"
)

var stdConn Client

var (
	DefaultPlayer = termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack}
	StairsUp      = termbox.Cell{'>', termbox.ColorDefault, termbox.ColorDefault}
	StairsDown    = termbox.Cell{'<', termbox.ColorDefault, termbox.ColorDefault}
)
