package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

// BorderSet is a set of characters that can be used to border a UI element.
// BorderSets must be laid out in the following order:
//
// Horizontal, Vertical
// Top-Left, Top-Middle, Top-Right
// Left-Middle, Center, Right-Middle
// Bottom-Left, Bottom-Middle, Bottom-Right
//
// TODO: Add remaining borders.
type BorderSet TileSet

const (
	LightBorder  BorderSet = "─│┌┬┐├┼┤└┴┘"
	HeavyBorder            = "━┃┏┳┓┣╋┫┗┻┛"
	DoubleBorder           = "═║╔╦╗╠╬╣╚╩╝"
)

// Border is a border around a UI element.
type Border struct {
	BorderSet // The runes to use for the border.
	Visible   bool
}

// Draw prints the border into termbox. Borders get drawn after the elements
// inside them.
func (b *Border) Draw(bounds Bounds) {
	if !b.Visible {
		return
	}
	defer termbox.Flush()

	// Top-Left corner.
	Ox, Oy := bounds[0].Ints()
	// Bottom-Right corner.
	w, h := bounds[1].Ints()

	s := []rune(fmt.Sprint(b.BorderSet))

	// Print the horizontals
	for x := Ox; x < w-1; x++ {
		termbox.SetCell(x, Oy, s[0], termbox.ColorDefault, termbox.ColorBlack)
		// termbox.SetCell(x, h/2, s[0], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h-1, s[0], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := Oy; y < h-1; y++ {
		termbox.SetCell(Ox, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
		// termbox.SetCell(w/2, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w-1, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(Ox, Oy, s[2], termbox.ColorDefault, termbox.ColorBlack)
	// termbox.SetCell(w/2, Oy, s[3], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, Oy, s[4], termbox.ColorDefault, termbox.ColorBlack)
	// termbox.SetCell(Ox, h/2, s[5], termbox.ColorDefault, termbox.ColorBlack)
	// termbox.SetCell(w/2, h/2, s[6], termbox.ColorDefault, termbox.ColorBlack)
	// termbox.SetCell(w-1, h/2, s[7], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h-1, s[8], termbox.ColorDefault, termbox.ColorBlack)
	// termbox.SetCell(w/2, h-1, s[9], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, s[10], termbox.ColorDefault, termbox.ColorBlack)
}

type TileSet string
