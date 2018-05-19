package ui

import (
	"bytes"
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"strings"
)

func CellsToByte(cells [][]termbox.Cell, x1, y1, x2, y2 int) []byte {
	var buff bytes.Buffer
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			buff.WriteRune(cells[x][y].Ch)
		}
		buff.WriteRune('\n')
	}
	return buff.Bytes()
}

// DrawAt draws the given interface in termbox starting from the given location (0x, 0y).
// Currently, this will overwrite any existing cells.
//
// DrawAt will attempt to convert v to a string, and fall back on fmt.Sprint(v).
//
// DrawAt returns an OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawAt(Ox, Oy int, v interface{}) error {
	var tmp string
	switch val := v.(type) {
	case []byte:
		tmp = string(val)
	case []rune:
		tmp = string(val)
	case []string:
		tmp = strings.Join(val, "\n")
	default:
		tmp = fmt.Sprint(val)
	}
	return DrawString(Ox, Oy, termbox.ColorDefault, termbox.ColorDefault, tmp)
}

// DrawRawString prints a string starting at the given coordinates (Ox, Oy).
// Line break ('\n') and carriage return ('\r') characters are not handled
// specially and will appear as spaces.
//
// DrawRawString returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawRawString(Ox, Oy int, fg, bg termbox.Attribute, v interface{}) error {
	defer termbox.Flush()
	x, y := Ox, Oy
	str := fmt.Sprint(v)
	for _, r := range str {
		termbox.SetCell(x, y, r, fg, bg)
		x++
	}
	return OutOfScreenBoundry(Bounds{engine.Point{Ox, Oy}, engine.Point{x, y}})
}

// DrawString prints a string starting at the given coordinates (Ox, Oy).
//
// Line Break ('\n') runes will move the "cursor" to the beginning of the next line,
//
// Carriage Return ('\r') runes will move the "cursor" to the beginning of the current line.
//
// DrawString returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func DrawString(Ox, Oy int, fg, bg termbox.Attribute, v interface{}) error {
	defer termbox.Flush()
	x, y := Ox, Oy
	str := fmt.Sprint(v)
	for _, r := range str {
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

// PrintScreen returns the current contents of termbox.
func PrintScreen() ([][]termbox.Cell, error) {
	w, h := termbox.Size()
	fmt.Println("w", w, "h", h)
	if w == 0 || h == 0 {
		return nil, errors.New("Termbox has no size. Has termbox been initialized?")
	}
	cells := termbox.CellBuffer()
	runes := [][]termbox.Cell{}
	for x := 0; x < w; x++ {
		runes = append(runes, []termbox.Cell{})
		for y := 0; y < h; y++ {
			runes[x] = append(runes[x], cells[(y*w)+x])
		}
	}
	return runes, nil
}

// SetCell is a wrapper for termbox.SetCell, which takes Cell attributes individually.
// SetCell will set the state of the given Cell in termbox.
func SetCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}
