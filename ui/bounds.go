package ui

import (
	"fmt"
	engine "github.com/sbrow/gorogue"
)

// Bounds hold the top left-most and bottom right-most points of a UIElement
type Bounds [2]engine.Point

func NewBounds(x1, y1, x2, y2 int) Bounds {
	return Bounds{engine.Point{x1, y1}, engine.Point{x2, y2}}
}

/*func (b *Bounds) Grow() {
	b[0].Shift(engine.NorthEast)
	b[1].Shift(engine.SouthEast)
}
*/
func (b *Bounds) Shift(d engine.Direction) {
	b[0].Shift(d)
	b[1].Shift(d)
}

func (b *Bounds) Shrink() {
	b[0].Shift(engine.SouthEast)
	b[1].Shift(engine.NorthWest)
}

func (b *Bounds) String() string {
	return fmt.Sprintf("[%s, %s]", b[0].String(), b[1].String())
}
