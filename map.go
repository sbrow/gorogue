package gorogue

import (
	termbox "github.com/nsf/termbox-go"
)

// Map is a set of 2 dimensional planes, containing objects. Each map with active
// connections will continue to Tick.
//
// TODO: Implement Map chunking, allow sections of larger maps to be evaluated independently,
// Evaluating their Tick before others, but pausing before beginning the next Tick.
type Map struct {
	Name   string
	Height int
	Width  int
	Actors Actors
}

func NewMap(w, h int, name string) *Map {
	m := &Map{}
	m.Width = w
	m.Height = h
	m.Name = name
	m.Actors = Actors([]Actor{})
	return m
}

func (m *Map) TileSlice(x1, y1, x2, y2 int) [][]termbox.Cell {
	ret := [][]termbox.Cell{}
	var i int

	// Draw Tiles
	for x := x1; x < x2; x++ {
		ret = append(ret, []termbox.Cell{})
		for y := y1; y < y2; y++ {
			ret[i] = append(ret[i], termbox.Cell{'.', termbox.ColorWhite,
				termbox.ColorBlack})
		}
		i++
	}

	// Draw Actors
	for _, a := range m.Actors {
		x, y := a.Pos().Ints()
		if x1 <= x && x <= x2 &&
			y1 <= y && y <= y2 {
			ret[x-x1][y-y1] = a.Sprite()
		}
	}
	return ret
}

// Tiles returns the same as calling TileSlice(0, 0, Map.Width, Map.Height)
func (m *Map) Tiles() [][]termbox.Cell {
	w, h := m.Width, m.Height
	return m.TileSlice(0, 0, w, h)
}
