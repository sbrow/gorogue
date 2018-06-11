package ui

import (
	"bytes"
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

// Cells returns the current contents of the UI.
// Returns an error if ternbox's width or height are zero.
func Cells() ([][]termbox.Cell, error) {
	termbox.Flush()
	w, h := Size()
	maxW, maxH := termbox.Size()
	if maxW == 0 || maxH == 0 {
		if !termbox.IsInit {
			return nil, errors.New("cells cannot be returned- termbox has not been initialized")
		}
		maxW = 80
		maxH = 24
		log.Println("Termbox has been initialized, but does not have size, assuming default (80x24).")
	}
	cells := termbox.CellBuffer()
	runes := [][]termbox.Cell{}
	for x := 0; x <= maxW; x++ {
		if x < w {
			runes = append(runes, []termbox.Cell{})
			for y := 0; y <= maxH; y++ {
				if y < h {
					runes[x] = append(runes[x], cells[(y*maxW)+x])
				}
			}
		}
	}
	return runes, nil
}

// Print draws the given interface in termbox starting from the Cell at (x, y).
// Print will overwrite any existing cells. v is interpreted using fmt.Print(v).
//
// Carriage return ('\r') runes will cause Print to continue printing from
// the start of the current line.
//
// Returns an OutOfScreenBoundryError if the drawing exceeds termbox's size.
func Print(x, y int, v ...interface{}) error {
	defer termbox.Flush()
	str := fmt.Sprint(v...)
	x1, y1 := x, y
	for _, r := range str {
		switch r {
		case '\n':
			y++
			fallthrough
		case '\r':
			x = x1
		default:
			SetCell(x, y, termbox.Cell{r, termbox.ColorDefault, termbox.ColorDefault})
			x++
		}
	}
	return outOfScreenBoundry(NewBounds(x1, y1, x, y))
}

// PrintRaw draws the given interface in termbox starting from the Cell at (x, y).
// PrintRaw will overwrite any existing cells. v is interpreted using fmt.Print(v).
//
// New Line ('\r') and Carriage Return ('\r') are interpreted as spaces.
//
// Returns an OutOfScreenBoundryError if the drawing exceeds termbox's size.
func PrintRaw(x, y int, v interface{}) error {
	defer termbox.Flush()
	x1, y1 := x, y
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	str := fmt.Sprint(v)
	for _, r := range str {
		SetCell(x, y, termbox.Cell{r, fg, bg})
		x++
	}
	return outOfScreenBoundry(NewBounds(x1, y1, x, y))
}

// outOfScreenBoundry determines whether the given boundaries are larger than
// the UI's current size. If they are, it returns an OutOfScreenBoundryError.
//
// Called by Draw and Print functions.
func outOfScreenBoundry(b Bounds) error {
	w, h := Size()
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
		return &OutOfScreenBoundryError{x, y, w, h}
	}
	return nil
}

// PrintScreen returns the contents of the UI.
// TODO: Move to RenderSystem
func PrintScreen() ([]byte, error) {
	var buff bytes.Buffer
	cells, err := Cells()
	if err != nil {
		return buff.Bytes(), err
	}
	for y := 0; y < len(cells[0]); y++ {
		for x := 0; x < len(cells); x++ {
			buff.WriteRune(cells[x][y].Ch)
		}
		buff.WriteRune('\n')
	}
	return buff.Bytes(), nil
}

type OutOfScreenBoundryError struct {
	X, Y int
	W, H int
}

func (o *OutOfScreenBoundryError) Error() string {
	return fmt.Sprintf("OutOfScreenBoundryError: point (%d, %d) "+
		"exceeds screen boundries [%d, %d]", o.X, o.Y, o.W, o.H)
}
