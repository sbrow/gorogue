package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
)

// Log is the standard Logger for gorogue.
var Log *log.Logger

var stdConn Client

// Sprites
var (
	DefaultPlayer Sprite = Sprite(termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack})
	StairsUp      Sprite = Sprite(termbox.Cell{'>', termbox.ColorDefault, termbox.ColorDefault})
	StairsDown    Sprite = Sprite(termbox.Cell{'<', termbox.ColorDefault, termbox.ColorDefault})
)

// Tiles
var (
	EmptyTile Tile = NewTile(termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack})
	FloorTile      = NewTile(termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack})
	WallTile       = NewTile(termbox.Cell{'#', termbox.ColorWhite, termbox.ColorBlack})
)
