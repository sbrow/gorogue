package example

import (
	engine "github.com/sbrow/gorogue"
	. "github.com/sbrow/gorogue/lib"
	"log"
)

// Map is a 2 dimensional plane containing tiles, objects and Actors. Each map will continue
// to Tick, so long as it has at least one active connection.
type Map struct {
	Name    string // The key that identifies this map in the server.
	Height  int    // The number of vertical tiles.
	Width   int    // The number of horizontal tiles.
	Players Actors
	Tiles   [][]engine.Tile
	ticks   int // The number of times this map has called Tick()
	actions chan int
	results chan bool
}

// NewMap creates a new, empty map with given dimensions and names.
func NewMap(w, h int, name string) *Map {
	m := &Map{}
	m.Width = w
	m.Height = h
	m.Name = name
	for x := 0; x < w; x++ {
		m.Tiles = append(m.Tiles, []engine.Tile{})
		for y := 0; y < h; y++ {
			m.Tiles[x] = append(m.Tiles[x], engine.Tile{FloorTile})
		}
	}
	m.actions = make(chan int)
	m.results = make(chan bool)
	return m
}

func (m *Map) Actors() []engine.Actor {
	a := []engine.Actor{}
	// for _, n := range m.NPCs {
	// a = append(a, Actor(n))
	// }
	for _, p := range m.Players {
		a = append(a, p)
	}
	return a
}

// Tick moves time forward one tick after it has received a valid Action from each
// Actor on the Map. Tick blocks all NPCs and connections.
//
// Currently, Actions are evaluated in FIFO order, meaning that Players' actions
//  will almost always be evaluated last.
func (m *Map) Tick() {
	queue := make([]int, len(m.Actors()))
	for i := 0; i < len(queue); i++ {
		queue[i] = <-m.actions
	}
	for _ = range queue {
		m.results <- true
	}
	m.ticks++
	log.Printf("Tick. (%d)\n=================================", m.ticks)
	m.Tick()
}

// TileSlice returns the contents of all tiles within the bounds of
// [(x1, y1), (x2, y2)]
func (m *Map) TileSlice(Ox, Oy, w, h int) [][]engine.Tile {
	ret := [][]engine.Tile{}
	x2, y2 := w, h
	if w > m.Width-1 {
		x2 = m.Width - 1
	}
	if h > m.Height-1 {
		y2 = m.Height - 1
	}

	// Draw Tiles
	i := 0
	for x := Ox; x <= x2; x++ {
		ret = append(ret, []engine.Tile{})
		for y := Oy; y <= y2; y++ {
			ret[i] = append(ret[i], m.Tiles[x][y])
		}
		i++
	}

	// Draw Actors
	for _, a := range m.Actors() {
		x, y, _ := a.Pos().Ints()
		if Ox <= x && x <= x2 &&
			Oy <= y && y <= y2 {
			ret[x-Ox][y-Oy] = engine.Tile{a.Sprite()}
		}
	}
	return ret
}

// AllTiles returns all of the map's tiles. It is congruent with
// calling TileSlice(0, 0, Map.Width, Map.Height)
func (m *Map) AllTiles() [][]engine.Tile {
	w, h := m.Width, m.Height
	return m.TileSlice(0, 0, w, h)
}

// WaitForTurn blocks an actor until Tick gives them priority.
func (m *Map) WaitForTurn(i int) bool {
	m.actions <- i
	return <-m.results
}
