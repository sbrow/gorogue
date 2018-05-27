package components

import (
	"engo.io/ecs"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

// BasicFace is an interface that allows BasicEntities to
// be added to systems automatically.
type BasicFace interface {
	GetBasicEntity() *ecs.BasicEntity
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

// PosFace is an interface that allows entities with a
// Pos component to be added to systems automatically.
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
