package gorogue

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

// Arguments passed from client to server
// when calling a Server.Move().
type MoveArgs struct {
	Actors Actors
	Points []Point
}

// Point represents a point in 2 dimensional space.
type Point struct {
	X int
	Y int
}

// Ints returns the point as a pair of ints.
func (p *Point) Ints() (x, y int) {
	return p.X, p.Y
}

// Response passed from server to client
// when calling  Server.Spawn().
type SpawnReply struct {
	Map    *string
	Actors Actors
}
