package ui

import (
	termbox "github.com/nsf/termbox-go"
	engine "github.com/sbrow/gorogue"
)

// View is a window into a Map. Views can be any size.
type View struct {
	ui     *UI
	border *Border
	bounds Bounds
	Origin engine.Point // Where this view is located in the UI.
	Map    *engine.Map  // The Map data is drawn from.
}

// NewView returns a newly created view.
//
// bounds is the portion of the map that
// you want displayed.
//
// origin is the location in the UI where
// you want this view to be placed.
func NewView(bounds Bounds, m *engine.Map, origin engine.Point) *View {
	v := &View{bounds: bounds, Origin: origin, Map: m}
	return v
}

func (v *View) Bounds() Bounds {
	return v.bounds
}

func (v *View) Border() *Border {
	return v.border
}

// Draw displays the view in termbox.
func (v *View) Draw() error {
	defer termbox.Flush()
	bounds := v.bounds
	origin := v.Origin
	if v.ui.border != nil {
		if v.ui.border.Visible {
			// TODO: Shift bounds accordingly
		}
	}

	// Get tiles from the map
	tiles := v.Map.TileSlice(bounds[0].X, bounds[0].Y, bounds[1].X,
		bounds[1].Y)

	DrawString(81, 0, termbox.ColorDefault, termbox.ColorDefault, bounds.String())
	DrawString(81, 1, termbox.ColorDefault, termbox.ColorDefault, len(tiles))
	DrawString(81, 2, termbox.ColorDefault, termbox.ColorDefault, len(tiles[0]))

	// Draw the tiles.
	var x, y int
	for y = 0; y < len(tiles[0]); y++ {
		for x = 0; x < len(tiles); x++ {
			SetCell(x+origin.X, y+origin.Y, termbox.Cell(tiles[x][y].Sprite))
		}
		for x = len(tiles); x <= bounds[1].X; x++ {
			SetCell(x+origin.X, y+origin.Y, engine.EmptyTile.Cell())
		}
	}
	for y = len(tiles[0]); y <= bounds[1].Y; y++ {
		for x = 0; x < bounds[1].X; x++ {
			SetCell(x+origin.X, y+origin.Y, engine.EmptyTile.Cell())
		}
	}
	return nil
}

func (v *View) SetBorder(b *Border) {
	v.border = b
}

func (v *View) SetUI(ui *UI) {
	v.ui = ui
}

func (v *View) Type() UIElementType {
	return UITypeView
}

func (v *View) UI() *UI {
	return v.ui
}
