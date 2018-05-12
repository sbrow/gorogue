package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
)

// Map is a 2 dimensional plane containing tiles, objects and Actors. Each map will continue
// to Tick, so long as it has at least one active connection.
type Map struct {
	Name    string            // The key that identifies this map in the server.
	Height  int               // The number of vertical tiles.
	Width   int               // The number of horizontal tiles.
	NPCs    []NPC             // Non-player characters.
	Players map[string]Player // Player characters
	ticks   int               // The number of times this map has called Tick()
	actions chan Actor
	results chan bool
}

// NewMap creates a new, empty map with given dimensions and names.
func NewMap(w, h int, name string) *Map {
	m := &Map{}
	m.Width = w
	m.Height = h
	m.Name = name
	m.Players = make(map[string]Player, 0)
	m.actions = make(chan Actor)
	m.results = make(chan bool)
	return m
}

func (m *Map) Actors() Actors {
	a := []Actor{}
	for _, n := range m.NPCs {
		a = append(a, Actor(n))
	}
	for _, p := range m.Players {
		a = append(a, Actor(p))
	}
	return Actors(a)
}

// Tick moves time forward one tick after it has received a valid Action from each
// Actor on the Map. Tick blocks all NPCs and connections.
//
// Currently, Actions are evaluated in FIFO order, meaning that Players' actions
//  will almost always be evaluated last.
func (m *Map) Tick() {
	queue := make(Actors, len(m.Actors()))
	for i := 0; i < len(queue); i++ {
		queue[i] = <-m.actions
	}
	for _ = range queue {
		m.results <- true
	}
	m.ticks++
	log.Printf("Tick. (%d)\n", m.ticks)
	m.Tick()
}

// TileSlice returns the contents of all tiles within the bounds of
// [(x1, y1), (x2, y2)]
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
	for _, a := range m.Actors() {
		x, y, _ := a.Pos().Ints()
		if x1 <= x && x <= x2 &&
			y1 <= y && y <= y2 {
			ret[x-x1][y-y1] = a.Sprite()
		}
	}
	return ret
}

// Tiles returns all of the map's tiles. It is congruent with
// calling TileSlice(0, 0, Map.Width, Map.Height)
func (m *Map) Tiles() [][]termbox.Cell {
	w, h := m.Width, m.Height
	return m.TileSlice(0, 0, w, h)
}

// WaitForTurn blocks an actor until Tick gives them priority.
func (m *Map) WaitForTurn(a Actor) bool {
	m.actions <- a
	return <-m.results
}
