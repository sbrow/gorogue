package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

// View is a window into a Map. Views can be any size.
type View struct {
	ui *UI
	// border *Border
	anchor engine.Point // Where this view is located in the UI.
	bounds Bounds
	origin engine.Point
	size   engine.Point
	Map    *engine.Map // The Map data is drawn from.
}

// NewView returns a newly created view.
//
// bounds is the portion of the map that
// you want displayed.
//
// anchor is the location in the UI where
// you want this view to be placed.
func NewView(mapBounds Bounds, m *engine.Map, anchor engine.Point) *View {
	v := &View{
		origin: mapBounds[0],
		size:   mapBounds[1],
		anchor: anchor,
		Map:    m,
	}
	return v
}

func (v *View) Bounds() Bounds {
	bounds := Bounds{v.anchor, v.anchor}
	uibounds := v.ui.InnerBounds()
	bounds[0].Add(uibounds[0])
	bounds[1].Add(v.size)

	if bounds[1].X > uibounds[1].X {
		bounds[1].X = uibounds[1].X
	}
	if bounds[1].Y > uibounds[1].Y {
		bounds[1].Y = uibounds[1].Y
	}
	DrawAt(81, 9, bounds)
	return bounds
}

/*func (v *View) Border() *Border {
	return v.border
}*/

// Draw displays the view in termbox.
func (v *View) Draw() error {
	defer termbox.Flush()
	bounds := v.Bounds()
	anchor := bounds[0]

	// Get tiles from the map
	tiles := v.Tiles()

	// Draw the tiles.
	var x, y int
	for y = 0; y < len(tiles[0]); y++ {
		for x = 0; x < len(tiles); x++ {
			SetCell(x+anchor.X, y+anchor.Y, termbox.Cell(tiles[x][y].Sprite))
		}
		for x = len(tiles); x <= bounds[1].X; x++ {
			// SetCell(x+anchor.X, y+anchor.Y, engine.EmptyTile.Cell())
		}
	}
	for y = len(tiles[0]); y <= bounds[1].Y; y++ {
		for x = 0; x < bounds[1].X; x++ {
			// SetCell(x+anchor.X, y+anchor.Y, engine.EmptyTile.Cell())
		}
	}
	return nil
}

func (v *View) Origin() engine.Point {
	return v.origin
}

/*func (v *View) SetBorder(b *Border) {
	v.border = b
}*/

func (v *View) SetUI(ui *UI) {
	v.ui = ui
}

func (v *View) Tiles() [][]engine.Tile {
	b := v.Bounds()
	x1, y1 := v.origin.X, v.origin.Y
	x2, y2 := b[1].Ints()
	x2 -= x1
	y2 -= y1
	DrawAt(81, 0, x1, y1, x2, y2)
	return v.Map.TileSlice(x1, y1, x2, y2)
}

func (v *View) Type() UIElementType {
	return UITypeView
}

func (v *View) UI() *UI {
	return v.ui
}
