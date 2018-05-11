package gorogue

type Action struct {
	Caller string
	Args   [][]byte
}

type ActionResponse struct {
	Msg   string
	Reply bool
}

// Direction represents the cardinal and ordinal directions.
// North points towards the top of the screen, east points to the right, etc.
//
// Bitwise operations are performed on Directions.
type Direction uint8

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
type Point struct {
	X int
	Y int
}

// Ints returns the point as a pair of ints.
func (p *Point) Ints() (x, y int) {
	return p.X, p.Y
}

// Pos represents the position of an object.
type Pos struct {
	Point
	Map int
}

func NewPos(x, y, z int) *Pos {
	return &Pos{Point{x, y}, z}
}

// Ints returns the point as an int triple.
func (p *Pos) Ints() (x, y, z int) {
	return p.X, p.Y, p.Map
}
