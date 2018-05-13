package lib

import termbox "github.com/nsf/termbox-go"

var (
	DefaultPlayer termbox.Cell = termbox.Cell{'@', termbox.ColorWhite, termbox.ColorBlack}
	EmptyTile                  = termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlack}
	FloorTile                  = termbox.Cell{'.', termbox.ColorWhite, termbox.ColorBlack}
)
