package gorogue

import (
	"errors"
	"log"
)

// Map is a 2 dime:nsional plane containing tiles, objects and Actors. Each map will continue
// to Tick, so long as it has at least one active connection.
type Map struct {
	Name    string // The key that identifies this map in the server.
	ID      uint8
	Height  int // The number of vertical tiles.
	Width   int // The number of horizontal tiles.
	Players map[string]Actor
	Tiles   [][]Tile
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
		m.Tiles = append(m.Tiles, []Tile{})
		for y := 0; y < h; y++ {
			m.Tiles[x] = append(m.Tiles[x], FloorTile)
		}
	}
	m.Players = map[string]Actor{}
	m.actions = make(chan int)
	m.results = make(chan bool)
	return m
}

func (m *Map) Actors() []Actor {
	a := []Actor{}
	for _, p := range m.Players {
		a = append(a, p)
	}
	return a
}

// AllTiles returns all of the map's tiles. It is congruent with
// calling TileSlice(0, 0, Map.Width, Map.Height)
func (m *Map) AllTiles() [][]Tile {
	w, h := m.Width, m.Height
	return m.TileSlice(0, 0, w, h)
}

func (m *Map) Move(a MoveAction) error {
	actor := a.Target.(Actor)

	// Assert that the Pos points to this Map.
	if a.Pos.Map != m.ID {
		return errors.New("Pointing to the wrong map.")
	}

	// Assert that the Pos is within the bounds of the Map.
	if a.Pos.X >= m.Width || a.Pos.Y >= m.Height ||
		a.Pos.X < 0 || a.Pos.Y < 0 {
		return errors.New("Can't move outside the map.")
	}

	// If all our assertions are correct, move the Actor.
	actor.SetPos(&a.Pos)
	return nil
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
// [x1, y1], [x2, y2]
func (m *Map) TileSlice(x1, y1, x2, y2 int) [][]Tile {
	ret := [][]Tile{}
	if x2 > m.Width-1 {
		x2 = m.Width - 1
	}
	if y2 > m.Height-1 {
		y2 = m.Height - 1
	}

	// Draw Tiles
	i := 0
	for x := x1; x <= x2; x++ {
		ret = append(ret, []Tile{})
		for y := y1; y <= y2; y++ {
			ret[i] = append(ret[i], m.Tiles[x][y])
		}
		i++
	}

	// Draw Actors
	for _, a := range m.Actors() {
		x, y, _ := a.Pos().Ints()
		if x1 <= x && x <= x2 &&
			y1 <= y && y <= y2 {
			ret[x-x1][y-y1] = Tile{Sprite: a.Sprite()}
		}
	}
	return ret
}

// WaitForTurn blocks an actor until Tick gives them priority.
func (m *Map) WaitForTurn(i int) bool {
	m.actions <- i
	return <-m.results
}
