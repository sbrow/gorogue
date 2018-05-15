package gorogue

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

var stdConn Client

var stdUI *UI

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
	return OutOfScreenBoundryError(Bounds{
		Point{Ox, Oy},
		Point{Ox + len(cells), Oy + len(cells[0])},
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
	return OutOfScreenBoundryError(Bounds{Point{Ox, Oy}, Point{x, y}})
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
	return OutOfScreenBoundryError(Bounds{Point{Ox, Oy}, Point{x, y}})
}

// Returned after an element is drawn.
// Returns nil if b does not exceed the terminal's size.
func OutOfScreenBoundryError(b Bounds) error {
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

// Border is a border around a UI element.
type Border struct {
	BorderSet // The runes to use for the border.
	UIElement // The element this is bordering.
}

// Bounds returns the bounds of b's UIElement.
func (b *Border) Bounds() Bounds {
	return b.UIElement.Bounds()
}

// Draw prints the border into termbox. Borders get drawn after the elements
// inside them.
func (b *Border) Draw() {
	defer termbox.Flush()
	bounds := b.Bounds()

	// Top-Left corner.
	Ox, Oy := bounds[0].Ints()
	// Bottom-Right corner.
	w, h := bounds[1].Ints()

	s := []rune(fmt.Sprint(b.BorderSet))

	// Print the horizontals
	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, Oy, s[0], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h-1, s[0], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := 1; y < h-1; y++ {
		termbox.SetCell(Ox, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w-1, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
	}

	// Print the corners.
	termbox.SetCell(Ox, Oy, s[2], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, Oy, s[4], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h-1, s[8], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, s[10], termbox.ColorDefault, termbox.ColorBlack)
}

// Type return's b's UIElementType.
func (b *Border) Type() UIElementType {
	return UITypeBorder
}

type TileSet string

// Bounds hold the top left-most and bottom right-most points of a UIElement
type Bounds [2]Point

// UI holds everything a player sees in game.
type UI struct {
	name   string
	bounds Bounds
	Border *Border          // The UI's border (if any).
	Views  map[string]*View // Views contained in this UI.
}

// New creates a new UI with a given name and size.
func NewUI(name string, w, h int) *UI {
	return &UI{
		name:   name,
		bounds: Bounds{Point{0, 0}, Point{w, h}},
		Border: nil,
		Views:  map[string]*View{},
	}
}

// AddView adds a view to this UI.
// The view is automatically adjusted to fit  if u has a Border.
func (u *UI) AddView(name string, v View) {
	u.Views[name] = &v
	V := u.Views[name]
	if u.Border != nil {
		V.Origin.X++
		V.Origin.Y++
	}
}

// Bounds return's u's bounds, including the border (if any).
func (u *UI) Bounds() Bounds {
	if u.Border == nil {
		return u.bounds
	}
	pt := u.bounds[0]
	pt2 := u.bounds[1]
	pt2.X += 2
	pt2.Y += 2
	return Bounds{pt, pt2}
}

// Draw displays the UI in termbox. UIElements are drawn in the following order:
//
// Views, Border.
func (u *UI) Draw() error {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	if err != nil {
		return err
	}
	stdConn.Ping()
	// Print each view
	for _, v := range u.Views {
		err := v.Draw()
		if err != nil {
			return err
		}
	}

	// Print the UI's border.
	if u.Border != nil {
		u.Border.Draw()
	}
	return nil
}

func (u *UI) Name() string {
	return u.name
}

// Run runs the active UI.
func (u *UI) Run() {
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)

	for {
		u.Draw()
		action, err := Input()
		if err != nil { //&& err.Error() != KeyNotBoundError { TODO: fix
			// panic(err)
		}
		if action != nil {
			err := stdConn.HandleAction(action)
			if err != nil {
				return
			}
		}
	}
}

func (u *UI) SetBorder(b BorderSet) {
	var delta int

	// If we don't already have a border, adjust our
	// bounds to fit.
	if u.Border == nil {
		delta = 1
	} else {
		u.Border = &Border{b, u}
		return
	}
	u.Border = &Border{b, u}
	for _, v := range u.Views {
		v.Origin.X += delta
		v.Origin.Y += delta
	}
}

func (u *UI) Type() UIElementType {
	return UITypeUI
}

// UIElement is anything that can show up in termbox,
// including UIs, Views and Borders
type UIElement interface {
	Name() string
	Type() UIElementType
	Bounds() Bounds
}

// UIElementType is an enum of valid UIElements.
type UIElementType uint8

const (
	UITypeBorder UIElementType = iota
	UITypeUI
	UITypeView
)

// View is a window into a Map. Views can be any size.
type View struct {
	Bounds
	Origin Point   // Where this view is located in the UI.
	Map    *string // The Map data is drawn from.
}

// NewView returns a newly created view.
//
// bounds is the portion of the map that
// you want displayed.
//
// origin is the location in the UI where
// you want this view to be placed.
func NewView(bounds Bounds, m *string, origin Point) *View {
	v := &View{Bounds: bounds, Origin: origin, Map: m}
	return v
}

// Draw displays the view in termbox.
func (v *View) Draw() error {
	defer termbox.Flush()

	// TODO: (10) Squad map get
	// FIXME:
	m := stdConn.Maps()[stdConn.Squad()[0].Pos().Map]
	// m := &Map{}

	// Get tiles from the map
	tiles := m.TileSlice(v.Bounds[0].X, v.Bounds[0].Y, v.Bounds[1].X,
		v.Bounds[1].Y)

	// Draw the tiles.
	var x, y int
	for y = 0; y < len(tiles[0]); y++ {
		for x = 0; x < len(tiles); x++ {
			cell := tiles[x][y].Sprite
			termbox.SetCell(x+v.Origin.X, y+v.Origin.Y, cell.Ch, cell.Fg, cell.Bg)
		}
		for x = len(tiles); x <= v.Bounds[1].X; x++ {
			SetCell(x+v.Origin.X, y+v.Origin.Y, EmptyTile.Cell())
		}
	}
	for y = len(tiles[0]); y <= v.Bounds[1].Y; y++ {
		for x = 0; x < v.Bounds[1].X; x++ {
			SetCell(x+v.Origin.X, y+v.Origin.Y, EmptyTile.Cell())
		}
	}
	return nil
}

func SetCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}
