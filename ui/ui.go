// Package ui is responsible for drawing the user interface and interpreting
// user input.
package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

var std *ui

// UI holds everything a player sees in game.
type ui struct {
	border   *Border              // The UI's border (if any).
	elements map[string]UIElement // Views contained in this UI.
	size     engine.Point
}

// New creates a new UI with a given size.
func New(w, h int) {
	std = &ui{}
	std.border = nil
	if w < 1 || h < 1 {
		std.size = engine.Point{-1, -1}
	} else {
		std.size = engine.Point{w - 1, h - 1}
	}
	std.elements = map[string]UIElement{}
}

// Add adds a UIElement to this UI.
func Add(name string, e UIElement) UIElement {
	e.SetBounds(*InnerBounds())
	std.elements[name] = e
	return std.elements[name]
}

// Draw displays the UI. UIElements are drawn in the following order:
//
// Views, Border.
func Draw() error {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return err
	}
	termbox.Flush()
	// Print each view
	for _, e := range std.elements {
		err := e.Draw()
		if err != nil {
			return err
		}
	}

	// Print the UI's border.
	if std.border != nil {
		std.border.Draw(*OuterBounds())
	}
	return nil
}

// Used for testing.
func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	Draw()
}

func InnerBounds() *Bounds {
	bounds := OuterBounds()
	if std.border != nil && std.border.Visible {
		bounds.Shrink()
	}
	return bounds
}

// OuterBounds return's u's bounds, including the border (if any).
func OuterBounds() *Bounds {
	var bounds Bounds
	// if !std.scales {
	bounds[1] = std.size
	// } else {
	// 		bounds[1] = engine.NewPoint(termbox.Size())
	// }
	return &bounds
}

// Run starts drawing the UI and accepting user input.
func Run() {
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)

	for {
		if err := Draw(); err != nil {
			panic(err)
		}
		action, err := engine.Input()
		if err != nil {
			engine.Log.Println("error: err")
		} else if action != nil {
			err := engine.HandleAction(action)
			if err != nil {
				if err.Error() == "Leaving..." { // TODO: Fix
					return
				}
				panic(err)
			}
		}
	}
}

// SetCell is a wrapper for termbox.SetCell, which takes Cell attributes individually.
// SetCell will set the state of the given Cell in termbox.
func SetCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}

// Size returns the size of the UI, in Cells.
func Size() (w, h int) {
	return OuterBounds().Size()
}

func SetBorder(set BorderSet, vis bool) {
	std.border = NewBorder(set, vis)
}

// UIElement is anything that can show up in termbox,
// including UIs, Views and Borders
type UIElement interface {
	// Border() *Border
	Bounds() Bounds
	Draw() error
	SetBounds(b Bounds)
	Type() UIElementType
}

// UIElementType is an enum of valid UIElements.
type UIElementType uint8

const (
	UITypeUI UIElementType = iota
	UITypeView

	// UITypeTextField
)
