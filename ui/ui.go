// Package ui  handles user input and output for the gorogue engine.
package ui

import (
	"engo.io/ecs"
	termbox "github.com/nsf/termbox-go"
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/components"
	"github.com/sbrow/gorogue/systems"
)

// std is the standard UI.
var std *ui

// ui holds everything a player sees in game.
type ui struct {
	border   *Border              // The UI's border (if any).
	elements map[string]UIElement // Views contained in this UI.
	size     Point
	world    *ecs.World
	renderer *systems.Render
}

// Add adds a UIElement to this UI.
func Add(name string, e UIElement) UIElement {
	e.SetBounds(*InnerBounds())
	ent := struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{e.Bounds()[0].X, e.Bounds()[0].Y},
		components.Sprite{*e.GetTiles()},
	}
	std.world.AddEntity(&ent)
	std.elements[name] = e
	return std.elements[name]
}

// Init sets up a new UI with the given size. Only one UI can run at any given time.
func Init(w, h int) {
	std = &ui{}
	std.border = nil
	if w < 1 || h < 1 {
		std.size = Point{-1, -1}
	} else {
		std.size = Point{w - 1, h - 1}
	}
	std.elements = map[string]UIElement{}
	std.world = &ecs.World{}
	var renderable *systems.Renderable
	std.world.AddSystemInterface(&systems.Render{}, renderable, nil)
	r, ok := std.world.Systems()[0].(*systems.Render)
	if !ok {
		panic("System is not renderer")
	}
	std.renderer = r
}

// Draw displays the UI. Borders are drawn after their contents.
func Draw() error {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}
	defer termbox.Flush()

	std.renderer.Update(0)
	if std.border != nil {
		std.border.Draw(*OuterBounds())
	}
	return nil
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
	// 		bounds[1] = NewPoint(termbox.Size())
	// }
	return &bounds
}

// Run starts drawing the UI and accepting user input.
func Run() error {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)

	for {
		if err := Draw(); err != nil {
			panic(err)
		}
		action, err := Input()
		if err != nil {
			Log.Println(err)
		} else if action != nil {
			err := HandleAction(action)
			if err != nil {
				if err.Error() == "Leaving..." { // TODO: Fix
					return nil
				}
				return err
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
	GetTiles() *[][]termbox.Cell
}

// UIElementType is an enum of valid UIElements.
type UIElementType uint8

const (
	UITypeUI UIElementType = iota
	UITypeView
	UITypeTextField
)
