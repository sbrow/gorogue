package ui

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

// DrawAt draws the given cells in termbox at the given location (0x, 0y).
// Currently, this will overwrite any existing cells.
//
// DrawAt returns an OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawAt(cells [][]termbox.Cell, Ox, Oy int) error {
	defer termbox.Flush()
	for y := 0; y < len(cells[0]); y++ {
		for x := 0; x < len(cells); x++ {
			SetCell(Ox+x, Oy+y, cells[x][y])
		}
	}
	return OutOfScreenBoundry(Bounds{
		engine.Point{Ox, Oy},
		engine.Point{Ox + len(cells), Oy + len(cells[0])},
	})
}

// DrawRawString prints a string starting at the given coordinates (Ox, Oy).
// Line break ('\n') and carriage return ('\r') characters are not handled
// specially and will appear as spaces.
//
// DrawRawString returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawRawString(Ox, Oy int, fg, bg termbox.Attribute, s string) error {
	defer termbox.Flush()
	x, y := Ox, Oy
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		x++
	}
	return OutOfScreenBoundry(Bounds{engine.Point{Ox, Oy}, engine.Point{x, y}})
}

// DrawString prints a string starting at the given coordinates (Ox, Oy).
//
// Line Break ('\n') runes will move the cursor to the beginning of the next line,
//
// Carriage Return ('\r') runes will move the cursor to the beginning of the current line.
//
// DrawString returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawString(Ox, Oy int, fg, bg termbox.Attribute, s string) error {
	defer termbox.Flush()
	x, y := Ox, Oy
	for _, r := range s {
		switch r {
		case '\n':
			y++
			fallthrough
		case '\r':
			x = Ox
		default:
			termbox.SetCell(x, y, r, fg, bg)
			x++
		}
	}
	return OutOfScreenBoundry(Bounds{engine.Point{Ox, Oy}, engine.Point{x, y}})
}

// OutOfScreenBoundry determines whether the given boundries are larger than
// termbox's current size. If they are, it returns an OutOfScreenBoundryError.
//
// Called by functions that draw to termbox.
func OutOfScreenBoundry(b Bounds) error {
	w, h := termbox.Size()
	var x, y int

	// If any coordinate is outside the screen, return an error.
	switch {
	case b[0].X < 0:
		fallthrough
	case b[0].Y < 0:
		x, y = b[0].X, b[0].Y
		w, h = 0, 0
	case b[1].X > w:
		fallthrough
	case b[1].Y > h:
		x, y = b[1].X, b[1].Y
	}
	if x != 0 || y != 0 {
		return errors.New(fmt.Sprintf("OutOfScreenBoundryError: point (%d, %d) "+
			"exceeds screen boundries [%d, %d]", x, y, w, h))
	}
	return nil
}

// SetCell is a wrapper for termbox.SetCell, which takes Cell attributes individually.
// SetCell will set the state of the given Cell in termbox.
func SetCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}
