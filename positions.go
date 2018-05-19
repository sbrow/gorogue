package gorogue

import "fmt"

// Direction represents the cardinal and ordinal directions.
// North points towards the top of the screen, east points to the right, etc.
//
// Converting between coordinates and Directions is often done with Bitwise operations,
// hence why they are not laid out in perfect sequence.
type Direction uint8

func (d Direction) Point() Point {
	var p Point
	if d&North == North {
		p.Y--
	}
	if d&South == South {
		p.Y++
	}
	if d&East == East {
		p.X++
	}
	if d&West == West {
		p.X--
	}
	return p
}

const (
	North     Direction = 1 + iota // 0001
	East                           // 0010
	NorthEast                      // 0011
	West                           // 0100
	NorthWest                      // 0101
	South     Direction = 8        // 1000
	SouthEast Direction = 10       // 1010
	SouthWest Direction = 12       // 1100
)

// Point represents a coordinate pair.
//
// Points are most commonly used to locate Tiles on a Map and Cells in termbox.
type Point struct {
	X int
	Y int
}

func (p *Point) Add(pt Point) {
	p.X += pt.X
	p.Y += pt.Y
}

func (p *Point) Shift(d Direction) {
	delta := d.Point()
	p.X += delta.X
	p.Y += delta.Y
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d %d)", p.X, p.Y)
}

// Ints returns the point as a pair of ints.
func (p Point) Ints() (x, y int) {
	return p.X, p.Y
}

// Pos represents the position of an Object in the World. Point holds their location
// in the Map, and Map holds the index of their Map in World.Maps.
type Pos struct {
	Point
	Map uint8
}

func NewPos(x, y, Map int) *Pos {
	return &Pos{Point{x, y}, uint8(Map)}
}

// Ints returns the position as an ordered triple.
func (p *Pos) Ints() (x, y, z int) {
	return p.X, p.Y, int(p.Map)
}
