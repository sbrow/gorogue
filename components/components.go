package components

import (
	"fmt"

	"engo.io/engo/common"
	termbox "github.com/nsf/termbox-go"
)

// Action is a request that some function be performed.
type Action struct {
	Name   string        // The type of Action to perform.
	Caller string        // The name of the object that requested this Action.
	Args   []interface{} // Any other relevant parameters.
}

// ActionFace allows Actions to be added to systems automatically.
type ActionFace interface {
	GetAction() *Action
}

// GetAction implements the ActionFace interface
func (a *Action) GetAction() *Action {
	return a
}

// BasicFace allows BasicEntities to be added to systems automatically.
type BasicFace interface {
	common.BasicFace
}

// Pos represents a point in the Scene.
type Pos struct {
	X, Y int
}

// GetPos implements the PosFace interface.
func (p *Pos) GetPos() *Pos {
	return p
}

// String returns the Pos in string form.
func (p *Pos) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

// Ints returns the Pos as a pair of int objects.
func (p *Pos) Ints() (x, y int) {
	return p.X, p.Y
}

// PosFace allows entities with a Pos component to be added to systems automatically.
type PosFace interface {
	GetPos() *Pos
}

// Sprite is a component that contains a collection of tiles.
type Sprite struct {
	Tiles [][]termbox.Cell
}

// GetSprite implements the SpriteFace interface.
func (s *Sprite) GetSprite() *Sprite {
	return s
}

// SpriteFace is an interface that allows entities with a
// Sprite component to be added to systems automatically.
type SpriteFace interface {
	GetSprite() *Sprite
}
