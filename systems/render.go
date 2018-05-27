package systems

import (
	"bytes"
	"engo.io/ecs"
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue/components"
)

// Render handles displaying entities to the user.
type Render struct {
	ents  []renderEnt
	world *ecs.World
}

// Add adds an entity to the Render system.
func (r *Render) Add(basic *ecs.BasicEntity, pos *components.Pos, sprite *components.Sprite) {
	r.ents = append(r.ents, renderEnt{basic, pos, sprite})
}

// AddByInterface adds an entity to the system, asserting that it implements Renderable.
// It will panic if the assertion fails.
func (r *Render) AddByInterface(o ecs.Identifier) {
	obj, ok := o.(Renderable)
	if !ok {
		panic(fmt.Sprintf("%s is not Renderable.", o))
	}
	r.Add(obj.GetBasicEntity(), obj.GetPos(), obj.GetSprite())
}

// Cells returns the current contents of termbox within the bounds [0, 0] x [w, h].
// Cells returns an error if ternbox's hasn't been initialized.
func (r *Render) Cells(w, h int) ([][]termbox.Cell, error) {
	termbox.Flush()
	maxW, maxH := termbox.Size()
	if maxW == 0 || maxH == 0 {
		return nil, errors.New("Termbox has no size. Has termbox been initialized?")
	}
	cells := termbox.CellBuffer()
	runes := [][]termbox.Cell{}
	for x := 0; x <= maxW; x++ {
		if x < w {
			runes = append(runes, []termbox.Cell{})
			for y := 0; y <= maxH; y++ {
				if y < h {
					runes[x] = append(runes[x], cells[(y*maxW)+x])
				}
			}
		}
	}
	return runes, nil
}

// New creates a new Render system.
func (r *Render) New(w *ecs.World) {
	r.world = w
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

// PrintScreen returns the contents of termbox in writable form. It
// produces the same results as calling []byte(r.Cells(w, h)), where w and h are
// termbox's size.
func (r *Render) PrintScreen() ([]byte, error) {
	var buff bytes.Buffer
	w, h := termbox.Size()
	cells, err := r.Cells(w, h)
	if err != nil {
		return buff.Bytes(), err
	}
	for y := 0; y < len(cells[0]); y++ {
		for x := 0; x < len(cells); x++ {
			buff.WriteRune(cells[x][y].Ch)
		}
		buff.WriteRune('\n')
	}
	return buff.Bytes(), nil
}

// Remove removes the entity from the system.
func (r *Render) Remove(basic ecs.BasicEntity) {
	i := -1
	for j, ent := range r.ents {
		if ent.ID() == basic.ID() {
			i = j
			break
		}
	}
	if i >= 0 {
		r.ents = append(r.ents[:i], r.ents[i+1:]...)
	}
}

// Update draws all entities in the system.
func (r *Render) Update(dt float32) {
	defer termbox.Flush()
	w, h := termbox.Size()

	for _, ent := range r.ents {
		x1, y1 := ent.Pos.Ints()
		tiles := ent.Sprite.Tiles
		x2 := x1 + len(tiles)
		y2 := y1 + len(tiles[0])
		if x2 > w+1 {
			x2 = w + 1
		}
		if y2 > h {
			y2 = h + 1
		}
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				c := tiles[x-x1][y-y1]
				termbox.SetCell(x, y, c.Ch, c.Fg, c.Bg)
			}
		}
	}
}

/*func Print(x, y int, v ...interface{}) error {
	defer termbox.Flush()
	str := fmt.Sprint(v...)
	x1, _ := x, y
	for _, r := range str {
		switch r {
		case '\n':
			y++
			fallthrough
		case '\r':
			x = x1
		default:
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
	}
	return nil
}*/

// renderEnt is an unexported struct for Renderable entities.
type renderEnt struct {
	*ecs.BasicEntity
	*components.Pos
	*components.Sprite
}

// Renderable is the interface that must be filled for an entity to be added
// to a Render system.
type Renderable interface {
	// ecs.BasicFace

	components.BasicFace
	components.PosFace
	components.SpriteFace
}
