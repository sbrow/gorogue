package gorogue

import (
	"container/heap"
	"errors"
	"log"
	"math/rand"
)

// Map is a 2 dimensional plane containing tiles, objects and Actors. Each map will continue
// to Tick, so long as it has at least one active connection. (Local or remote).
type Map struct {
	Index   uint8            // The index of this map in its World.
	Height  int              // The number of vertical tiles.
	Width   int              // The number of horizontal tiles.
	Players map[string]Actor // Player characters on this Map.
	Tiles   [][]Tile         // The Tiles that make this map up.
	Ready   chan *Item
	World   World
	pq      priorityQueue
	ticks   int
}

// NewMap creates a new, empty map with the given dimensions.
func NewMap(w, h int) *Map {
	m := &Map{}
	m.Width = w
	m.Height = h
	m.Ready = make(chan *Item)
	m.pq = priorityQueue{}
	for x := 0; x < w; x++ {
		m.Tiles = append(m.Tiles, []Tile{})
		for y := 0; y < h; y++ {
			m.Tiles[x] = append(m.Tiles[x], FloorTile)
		}
	}
	m.Players = map[string]Actor{}
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
	if a.Pos.Map != m.Index {
		return errors.New("Pointing to the wrong map.")
	}

	// Assert that the Pos is within the bounds of the Map.
	if a.Pos.X >= m.Width || a.Pos.Y >= m.Height ||
		a.Pos.X < 0 || a.Pos.Y < 0 {
		return errors.New("Can't move outside the map.")
	}

	// If all our assertions are correct, move the Actor.
	actor.Ready()
	if err := actor.Move(&a.Pos); err != nil {
		return err
	}
	actor.Done()
	return nil
}

func (m *Map) Remove(str string) {
	delete(m.Players, str)
}

func (m *Map) Spawn(sa *SpawnAction) error {
	x := rand.Intn(m.Width)
	y := rand.Intn(m.Height)
	actor := m.World.Players()[sa.Caller]
	m.Players["Player_1"] = actor
	actor.SetMap(m)
	actor.SetPt(NewPt(x, y))
	if len(m.Players) == 1 {
		go m.Tick()
	}
	return nil
}

// Tick moves time forward one tick after it has received a valid Action from each
// Actor on the Map. Tick blocks all NPCs and connections.
//
// Currently, Actions are evaluated in FIFO order, meaning that Players' actions
// will almost always be evaluated last.
func (m *Map) Tick() {
	if len(m.Actors()) < 1 {
		return
	}
	heap.Init(&m.pq)
	for m.pq.Len() < len(m.Actors()) {
		heap.Push(&m.pq, <-m.Ready)
	}
	for m.pq.Len() > 0 {
		i := heap.Pop(&m.pq).(*Item)
		log.Println(i)
		i.Ch <- "It's your turn!"
		<-i.Ch
	}
	m.ticks++
	log.Printf("Tick! (%d)\n=========================\n", m.ticks)
	m.Tick()
}

// TileSlice returns the contents of all tiles within the bounds of
// [x1, y1], [x2, y2]. TileSlice's paramaters do not need to be within the
// Map's bounds, tiles outside will appear as BlankTiles
//
// TODO: Describe better.
func (m *Map) TileSlice(x1, y1, x2, y2 int) [][]Tile {
	ret := [][]Tile{}

	i := 0
	for x := x1; x <= x2; x++ {
		ret = append(ret, []Tile{})
		for y := y1; y <= y2; y++ {
			if x < 0 || x > m.Width-1 || y < 0 || y > m.Height-1 {
				ret[i] = append(ret[i], BlankTile)
			} else {
				ret[i] = append(ret[i], m.Tiles[x][y])
			}
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
