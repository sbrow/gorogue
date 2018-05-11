package gorogue

import (
	termbox "github.com/nsf/termbox-go"
	"log"
)

// Map 2 dimensional plane containing tiles, objects and Actors. Each map with active
// connections will continue to Tick.
//
// TODO: Implement Map chunking, allow sections of larger maps to be evaluated independently,
// Evaluating their Tick before others, but pausing before beginning the next Tick.
type Map struct {
	Name    string
	Height  int
	Width   int
	NPCs    Actors
	Players Actors // TODO: Temporary Solution
	ticks   int
	actions chan Actor
	results chan bool
}

// NewMap creates a new, empty map of given dimensions
func NewMap(w, h int, name string) *Map {
	m := &Map{}
	m.Width = w
	m.Height = h
	m.Name = name
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

// TileSlice returns the contents of all tiles within the bounds of
// (x1, y1) X (x2, y2).
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

// 1. Sender sends action to server
// 2. Server determines whether action is valid.
// 3. If it isn't, server returns failed action, then GOTO 1.
// 4. If it is, server adds action to queue. (BLOCKING)

// Tick collects actions from all actors on the map, and inserts them into an
// array at a random index. When tick has received an action from all Actors,
// it executes the actions in the array (in order), and then increments m.ticks and
// restarts.
func (m *Map) Tick() {
	queue := make(Actors, len(m.Actors()))
	// log.Println("Collecting Actions for next Tick...")
	for i := 0; i < len(queue); i++ {
		queue[i] = <-m.actions
	}
	// log.Println("Actions collected, sending responses...")
	for _ = range queue {
		m.results <- true
	}
	// log.Println("Reponses sent.")
	m.ticks++
	log.Printf("Tick. (%d)\n", m.ticks)
	m.Tick()
}

func (m *Map) WaitForTurn(a Actor) bool {
	m.actions <- a
	return <-m.results
}
