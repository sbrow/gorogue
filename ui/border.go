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
	HeavyBorder  BorderSet = "━┃┏┳┓┣╋┫┗┻┛"
	DoubleBorder           = "═║╔╦╗╠╬╣╚╩╝"
)

// TODO: Document
type BorderRune uint8

const (
	Horizontal BorderRune = iota
	Vertical
	LeftTop
	MiddleTop
	RightTop
	LeftMiddle
	Center
	RightMiddle
	LeftBottom
	MiddleBottom
	RightBottom
)

// Border is a border around a UI element.
type Border struct {
	BorderSet // The runes to use for the border.
	Visible   bool
}

func NewBorder(set BorderSet, vis bool) *Border {
	return &Border{
		BorderSet: set,
		Visible:   vis,
	}
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
		termbox.SetCell(x, Oy, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h-1, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := Oy; y < h-1; y++ {
		termbox.SetCell(Ox, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w-1, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(Ox, Oy, s[LeftTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, Oy, s[RightTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h-1, s[LeftBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, s[RightBottom], termbox.ColorDefault, termbox.ColorBlack)
}

type TileSet string
