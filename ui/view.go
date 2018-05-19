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
	return bounds
}

/*func (v *View) Border() *Border {
	return v.border
}*/

func (v *View) Center() engine.Point {
	return engine.Point{v.Width()/2 - 1, v.Height()/2 - 1}
}

// origin = Center - {w/2, h/2}

func (v *View) CenterView(p engine.Point) {
	p.Sub(v.Center())
	DrawAt(83, 5, p)
	v.origin = p
}

// Draw displays the view in termbox.
func (v *View) Draw() error {
	defer termbox.Flush()
	bounds := v.Bounds()
	anchor := bounds[0]
	termbox.SetCursor(v.Width()/2, v.Height()/2)

	// Get tiles from the map
	tiles := v.Tiles()

	// Draw the tiles.
	var x, y int
	for y = 0; y < len(tiles[0]); y++ {
		for x = 0; x < len(tiles); x++ {
			SetCell(x+anchor.X, y+anchor.Y, termbox.Cell(tiles[x][y].Sprite))
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
	x2 += x1 - 1
	y2 += y1 - 1
	DrawAt(85, 0, x1, y1, x2, y2)
	DrawAt(85, 1, " or: ", v.origin, " w: ", v.Width()/2, " h: ", v.Height()/2)
	DrawAt(85, 2, b)
	return v.Map.TileSlice(x1, y1, x2, y2)
}

func (v *View) Type() UIElementType {
	return UITypeView
}

func (v *View) UI() *UI {
	return v.ui
}

func (v *View) Width() int {
	b := v.Bounds()
	return b[1].X - b[0].X + 1
}

func (v *View) Height() int {
	b := v.Bounds()
	return b[1].Y - b[0].Y + 1
}
