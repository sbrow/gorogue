package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
)

// Log is the standard Logger for gorogue.
var Log *log.Logger

var stdConn Client

var (
	DefaultPlayer termbox.Cell = termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack}
	StairsUp                   = termbox.Cell{'>', termbox.ColorDefault, termbox.ColorDefault}
	StairsDown                 = termbox.Cell{'<', termbox.ColorDefault, termbox.ColorDefault}
)

var (
	EmptyTile Tile = NewTile(termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack})
	FloorTile Tile = NewTile(termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack})
	WallTile  Tile = NewTile(termbox.Cell{'#', termbox.ColorWhite, termbox.ColorBlack})
)
