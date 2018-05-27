package ui

import (
	termbox "github.com/nsf/termbox-go"
	. "github.com/sbrow/gorogue"
)

// View is a window into a Map. Views can be any size.
type View struct {
	// border *Border
	bounds Bounds
	origin Point
	Tiles  *[][]Tile
	cells  *[][]termbox.Cell
}

// NewView returns a newly created view.
//
// bounds is the portion of the map that
// you want displayed.
//
// anchor is the location in the UI where
// you want this view to be placed.
func NewView(origin Point, t *[][]Tile) *View {
	v := &View{
		origin: origin,
		Tiles:  t,
	}
	return v
}

func (v *View) Bounds() Bounds {
	return v.bounds
}

/*func (v *View) Border() *Border {
	return v.border
}*/

func (v *View) Center() Point {
	return Point{v.Width()/2 - 1, v.Height()/2 - 1}
	// return Point{v.Width() / 2, v.Height() / 2}
}

func (v *View) CenterView(p Point) {
	p.Sub(v.Center())
	v.origin = p
}

// Draw displays the view in termbox.
func (v *View) Draw() error {
	defer termbox.Flush()
	// bounds := v.Bounds()
	anchor := v.bounds[0]

	// Draw the tiles.
	var x, y int
	for y = 0; y < len((*v.Tiles)[0]); y++ {
		for x = 0; x < len(*v.Tiles); x++ {
			if x < 0 || x > v.Width()-1 || y < 0 || y > v.Height()-1 {
				SetCell(x+anchor.X, y+anchor.Y, BlankTile.Sprite)
			} else {
				SetCell(x+anchor.X, y+anchor.Y, termbox.Cell((*v.Tiles)[x][y].Sprite))
			}
		}
	}
	return nil
}

func (v *View) Origin() Point {
	return v.origin
}

/*func (v *View) SetBorder(b *Border) {
	v.border = b
}*/

func (v *View) SetBounds(b Bounds) {
	v.bounds = b
}

func (v *View) Size() (w, h int) {
	return v.Bounds().Size()
}

func (v *View) GetTiles() *[][]termbox.Cell {
	out := [][]termbox.Cell{}
	tiles := *v.Tiles
	i := 0
	for x := 0; x < len(tiles); x++ {
		out = append(out, []termbox.Cell{})
		for y := 0; y < len(tiles[0]); y++ {
			out[i] = append(out[i], tiles[x][y].Sprite)
		}
		i++
	}
	v.cells = &out
	return v.cells
}

func (v *View) Type() UIElementType {
	return UITypeView
}

func (v *View) Width() int {
	b := v.Bounds()
	return b[1].X - b[0].X + 1
}

func (v *View) Height() int {
	b := v.Bounds()
	return b[1].Y - b[0].Y + 1
}
