package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
)

// Log is the standard Logger for gorogue.
var Log *log.Logger

var StdConn Client

var (
	DefaultPlayer = termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack}
	StairsUp      = termbox.Cell{'>', termbox.ColorDefault, termbox.ColorDefault}
	StairsDown    = termbox.Cell{'<', termbox.ColorDefault, termbox.ColorDefault}
)
