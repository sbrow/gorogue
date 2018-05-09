package gorogue

// Direction represents the cardinal and ordinal directions.
// North points towards the top of the screen, east points to the right, etc.
type Direction uint8

const (
	North Direction = iota
	East
	South
	West
	NorthEast
	NorthWest
	SouthEast
	SouthWest
)

type MoveArgs struct {
	Actors
	Points []Point
}

type Point struct {
	X int
	Y int
}

func (p *Point) Ints() (x, y int) {
	return p.X, p.Y
}

type SpawnReply struct {
	Map    *string
	Actors Actors
}
