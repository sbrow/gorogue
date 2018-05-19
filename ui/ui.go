package ui

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

// Bounds hold the top left-most and bottom right-most points of a UIElement
type Bounds [2]engine.Point

func NewBounds(x1, y1, x2, y2 int) Bounds {
	return Bounds{engine.Point{x1, y1}, engine.Point{x2, y2}}
}

func (b *Bounds) Grow() {
	b[0].X--
	b[0].Y--
	b[1].X++
	b[1].Y++
}

func (b *Bounds) Shrink() {
	b[0].X++
	b[0].Y++
	b[1].X--
	b[1].Y--
}

func (b *Bounds) String() string {
	return fmt.Sprintf("[%s, %s]", b[0].String(), b[1].String())
}

// UI holds everything a player sees in game.
type UI struct {
	name     string
	client   engine.Client
	bounds   Bounds
	border   *Border              // The UI's border (if any).
	Elements map[string]UIElement // Views contained in this UI.
}

// New creates a new UI with a given name and size.
func NewUI(client engine.Client, name string, w, h int) *UI {
	return &UI{
		name:     name,
		client:   client,
		bounds:   Bounds{engine.Point{0, 0}, engine.Point{w, h}},
		border:   nil,
		Elements: map[string]UIElement{},
	}
}

// Add adds a UIElement to this UI.
func (u *UI) Add(name string, e UIElement) UIElement {
	u.Elements[name] = e
	e.SetUI(u)
	return u.Elements[name]
}

// Bounds return's u's bounds, including the border (if any).
func (u *UI) Bounds() Bounds {
	return u.bounds
}

// Draw displays the UI in termbox. UIElements are drawn in the following order:
//
// Views, Border.
func (u *UI) Draw() error {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return err
	}
	termbox.Flush()
	// Print each view
	for _, e := range u.Elements {
		err := e.Draw()
		if err != nil {
			return err
		}
	}

	// Print the UI's border.
	if u.border != nil {
		u.border.Draw(u.Bounds())
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
		action, err := engine.Input()
		if err != nil { //&& err.Error() != KeyNotBoundError { TODO: fix
			// panic(err)
		}
		if action != nil {
			err := u.client.HandleAction(action)
			if err != nil {
				return
			}
		}
	}
}

func (u *UI) SetBorder(b *Border) {
	u.border = b
}

func (u *UI) Type() UIElementType {
	return UITypeUI
}

// UIElement is anything that can show up in termbox,
// including UIs, Views and Borders
type UIElement interface {
	Border() *Border
	Bounds() Bounds
	Draw() error
	SetUI(u *UI)
	Type() UIElementType
	UI() *UI
}

// UIElementType is an enum of valid UIElements.
type UIElementType uint8

const (
	UITypeUI UIElementType = iota
	UITypeView
	UITypeTextField
)
