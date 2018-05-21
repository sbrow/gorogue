package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
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
//
// TODO: Test
type BorderSet engine.TileSet

const (
	LightBorder      BorderSet = "─│┌┬┐├┼┤└┴┘"
	HeavyLightBorder BorderSet = "━│┍┯┑┝┿┥┕┷┙"
	LightHeavyBorder BorderSet = "─┃┎┰┒┠╂┨┖┸┚"
	HeavyBorder      BorderSet = "━┃┏┳┓┣╋┫┗┻┛"
	DoubleBorder     BorderSet = "═║╔╦╗╠╬╣╚╩╝"
)

// TODO: Document
//
// TODO: Test
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
	BorderSet      // The RuneSet to use for the border.
	Visible   bool // Whether to draw the border.
}

// NewBorder returns a new Border with the given values.
func NewBorder(set BorderSet, vis bool) *Border {
	return &Border{
		BorderSet: set,
		Visible:   vis,
	}
}

// Draw prints the border in termbox. Borders get drawn after the elements
// inside them.
func (b *Border) Draw(bounds Bounds) {
	if !b.Visible {
		return
	}
	defer termbox.Flush()

	// Top-Left corner.
	x1, y1 := bounds[0].Ints()
	// Bottom-Right corner.
	x2, y2 := bounds[1].Ints()

	s := []rune(string(b.BorderSet))

	// Print the horizontals
	for x := x1; x < x2; x++ {
		termbox.SetCell(x, y1, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, y2, s[Horizontal], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := y1; y < y2; y++ {
		termbox.SetCell(x1, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x2, y, s[Vertical], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(x1, y1, s[LeftTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y1, s[RightTop], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x1, y2, s[LeftBottom], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(x2, y2, s[RightBottom], termbox.ColorDefault, termbox.ColorBlack)
}
