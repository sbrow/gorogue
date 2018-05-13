package client

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/sbrow/gorogue"
)

// DrawAt draws the given cells in termbox at the given location (0x, 0y).
// Currently, this will write over any existing cells.
//
// DrawAt returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func (u UI) DrawAt(cells [][]termbox.Cell, Ox, Oy int) error {
	defer termbox.Flush()
	for y := 0; y < len(cells[0]); y++ {
		for x := 0; x < len(cells); x++ {
			cell := cells[x][y]
			termbox.SetCell(Ox+x, Oy+y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
	return u.OutOfScreenBoundryError(Bounds{Point{Ox, Oy},
		Point{Ox + len(cells), Oy + len(cells[0])}})
}

// Returned after an element is drawn.
// Returns nil if b does not exceed the terminal's size.
func (u UI) OutOfScreenBoundryError(b Bounds) error {
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

// String prints an unterminated string, starting at the given coordinates (Ox, Oy).
//
// String returns OutOfScreenBoundryError if the drawing exceeds termbox's size.
func (u UI) String(Ox, Oy int, fg, bg termbox.Attribute, s string) error {
	defer termbox.Flush()
	x, y := Ox, Oy
	for _, c := range s {
		switch c {
		case '\r':
		case '\n':
			x--
			y++
			fallthrough
		default:
			termbox.SetCell(x, y, c, fg, bg)
			x++
		}
	}
	return u.OutOfScreenBoundryError(Bounds{Point{Ox, Oy}, Point{x, y}})
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

	// Print the corners.
	termbox.SetCell(Ox, Oy, s[0], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, Oy, s[1], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(Ox, h-1, s[2], termbox.ColorDefault, termbox.ColorBlack)
	termbox.SetCell(w-1, h-1, s[3], termbox.ColorDefault, termbox.ColorBlack)

	// Print the horizontals
	for x := 1; x < w-1; x++ {
		termbox.SetCell(x, Oy, s[4], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(x, h-1, s[4], termbox.ColorDefault, termbox.ColorBlack)
	}
	// Print the verticals
	for y := 1; y < h-1; y++ {
		termbox.SetCell(Ox, y, s[5], termbox.ColorDefault, termbox.ColorBlack)
		termbox.SetCell(w-1, y, s[5], termbox.ColorDefault, termbox.ColorBlack)
	}
}

// Type return's b's UIElementType.
func (b *Border) Type() UIElementType {
	return UITypeBorder
}

// BorderSet is a set of characters that can be used to border a UI element.
// BorderSets must be laid out in the following order:
//
// Top Left, Top Right, Bottom Left, Bottom Right, Horizontal, Vertical
// VerticalRight, VerticalLeft, LeftUp, Center,DownHorizontal
//
// TODO: Add remaining borders.
type BorderSet TileSet

type TileSet string

const (
	LightBorder  BorderSet = "┌┐└┘─│┤├┴┼┬"
	HeavyBorder            = "┏┓┗┛━┃┫┣┻╋┳"
	DoubleBorder           = "╔╗╚╝═║╣╠╩╬╦"
)

// Bounds hold the top left-most and bottom right-most points of a UIElement
type Bounds [2]Point

// UI holds everything a player sees in game.
//
// TODO: Improve UI description.
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
	Ping()
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
func Run() {
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)

main:
	for {
		ui.Draw()
		action, err := Input()
		if err != nil && err.Error() != KeyNotBoundError {
			panic(err)
		}
		if action != nil {
			switch action.Name {
			case "Quit":
				break main
			case "Move":
				Move(action)
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
	m := std.Maps[std.Squad[0].Pos().Map]

	// Get tiles from the map
	tiles := m.TileSlice(v.Bounds[0].X, v.Bounds[0].Y, v.Bounds[1].X,
		v.Bounds[1].Y)

	// Draw the tiles.
	for y := 0; y < len(tiles[0]); y++ {
		for x := 0; x < len(tiles); x++ {
			cell := tiles[x][y]
			termbox.SetCell(x+v.Origin.X, y+v.Origin.Y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
	return nil
}
