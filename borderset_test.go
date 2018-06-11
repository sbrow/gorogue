package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"testing"
)

func TestBorder(t *testing.T) {
	termbox.Init()
	defer termbox.Close()

	// Top-Left corner.
	Ox, Oy := 0, 0
	// Bottom-Right corner.
	w, h := 19, 11

	s := []rune(LightBorder)

	// Print the horizontals
	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, Oy, s[0], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h/2, s[0], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h-1, s[0], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := 1; y < h-1; y++ {
		termbox.SetCell(Ox, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w/2, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w-1, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(Ox, Oy, s[2], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w/2, Oy, s[3], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, Oy, s[4], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h/2, s[5], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w/2, h/2, s[6], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h/2, s[7], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h-1, s[8], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w/2, h-1, s[9], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, s[10], termbox.ColorDefault, termbox.ColorBlack)

	termbox.Flush()
}
