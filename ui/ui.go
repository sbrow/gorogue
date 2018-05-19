// Package ui is responsible for
package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

// UI holds everything a player sees in game.
type UI struct {
	name     string
	client   engine.Client
	bounds   Bounds
	border   *Border              // The UI's border (if any).
	Elements map[string]UIElement // Views contained in this UI.
}

// New creates a new UI with a given name and size.
func NewUI(client engine.Client, name string, x, y, w, h int) *UI {
	ui := &UI{}
	ui.client = client
	ui.border = nil
	ui.bounds[0] = engine.Point{x, y}
	ui.bounds[1] = engine.Point{x + w - 1, y + h - 1}
	ui.name = name
	ui.Elements = map[string]UIElement{}
	return ui
}

// Add adds a UIElement to this UI.
func (u *UI) Add(name string, e UIElement) UIElement {
	u.Elements[name] = e
	e.SetUI(u)
	return u.Elements[name]
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
		u.border.Draw(u.OuterBounds())
	}
	return nil
}

func (u *UI) InnerBounds() Bounds {
	bounds := u.bounds
	if u.border.Visible {
		bounds.Shrink()
	}
	return bounds
}

func (u *UI) Name() string {
	return u.name
}

// OuterBounds return's u's bounds, including the border (if any).
func (u *UI) OuterBounds() Bounds {
	return u.bounds
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
		if err != nil {
			engine.Log.Println("error:", err)
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
	// Border() *Border
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
