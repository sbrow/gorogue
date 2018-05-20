package ui

import (
	"bytes"
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

// Cells returns the current contents of the UI.
func Cells() ([][]termbox.Cell, error) {
	termbox.Flush()
	w, h := Size()
	w2, h2 := termbox.Size()
	if w2 == 0 || h2 == 0 {
		return nil, errors.New("Termbox has no size. Has termbox been initialized?")
	}
	cells := termbox.CellBuffer()
	runes := [][]termbox.Cell{}
	for x := 0; x <= w2; x++ {
		if x < w {
			runes = append(runes, []termbox.Cell{})
			for y := 0; y <= h2; y++ {
				if y < h {
					runes[x] = append(runes[x], cells[(y*w2)+x])
				}
			}
		}
	}
	return runes, nil
}

// Print draws the given interface in termbox starting from the given location (0x, 0y).
// Currently, this will overwrite any existing cells.
//
// Print will attempt to convert v to a string, and fall back on fmt.Sprint(v).
//
// Print returns an OutOfScreenBoundryError if the drawing exceeds termbox's size.
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
	return OutOfScreenBoundry(NewBounds(x1, y1, x, y))
}

// PrintRaw prints a string starting at the given coordinates (x, y).
// Line break ('\n') and carriage return ('\r') characters are not handled
// specially and will appear as spaces.
//
// PrintRaw returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func PrintRaw(x, y int, v interface{}) error {
	defer termbox.Flush()
	x1, y1 := x, y
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	str := fmt.Sprint(v)
	for _, r := range str {
		SetCell(x, y, termbox.Cell{r, fg, bg})
		x++
	}
	return OutOfScreenBoundry(NewBounds(x1, y1, x, y))
}

// OutOfScreenBoundry determines whether the given boundries are larger than
// termbox's current size. If they are, it returns an OutOfScreenBoundryError.
//
// Called by functions that draw to termbox.
func OutOfScreenBoundry(b Bounds) error {
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
		return errors.New(fmt.Sprintf("OutOfScreenBoundryError: point (%d, %d) "+
			"exceeds screen boundries [%d, %d]", x, y, w, h))
	}
	return nil
}

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
