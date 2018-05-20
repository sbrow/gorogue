// Package ui is responsible for drawing the user interface and interpreting
// user input.
package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
	"log"
)

var std *ui

// UI holds everything a player sees in game.
type ui struct {
	client   engine.Client
	bounds   Bounds
	border   *Border              // The UI's border (if any).
	elements map[string]UIElement // Views contained in this UI.
}

// New creates a new UI with a given size.
func New(client engine.Client, w, h int) {
	std = &ui{}
	std.client = client
	std.border = nil
	std.bounds[0] = engine.Point{0, 0}
	std.bounds[1] = engine.Point{w - 1, h - 1}
	std.elements = map[string]UIElement{}
}

// Add adds a UIElement to this UI.
func Add(name string, e UIElement) UIElement {
	std.elements[name] = e
	return std.elements[name]
}

// Draw displays the UI in termbox. UIElements are drawn in the following order:
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
		std.border.Draw(OuterBounds())
	}
	return nil
}

func Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	Draw()
}

func InnerBounds() Bounds {
	bounds := std.bounds
	if std.border.Visible {
		bounds.Shrink()
	}
	return bounds
}

// OuterBounds return's u's bounds, including the border (if any).
func OuterBounds() Bounds {
	return std.bounds
}

// Run runs the active UI.
func Run() {
	err := termbox.Init()
	defer termbox.Close()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)

	for {
		Draw()
		action, err := engine.Input()
		if err != nil {
			// engine.Log.Println("error:", err)
			log.Println("error: ", err)
		} else if action != nil {
			if std.client == nil {
				return
			}
			err := std.client.HandleAction(action)
			if err != nil {
				return
			}
		}
	}
}

// SetCell is a wrapper for termbox.SetCell, which takes Cell attributes individually.
// SetCell will set the state of the given Cell in termbox.
func SetCell(x, y int, c termbox.Cell) {
	termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
}

func Size() (w, h int) {
	return std.bounds.Size()
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
	Type() UIElementType
}

// UIElementType is an enum of valid UIElements.
type UIElementType uint8

const (
	UITypeUI UIElementType = iota
	UITypeView
	UITypeTextField
)
