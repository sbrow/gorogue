package gorogue

import (
	"encoding/json"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

type Object interface {
	Name() string
	ID() string
	Index() int
	MarshalJSON() ([]byte, error)
	Pos() *Pos
	SetIndex(i int)
	SetPos(p Pos)
	Sprite() termbox.Cell
	UnmarshalJSON(data []byte) error
}

type object struct {
	name   string
	index  int
	pos    *Pos
	sprite termbox.Cell
}

func newObject(name string, index int, pos *Pos, sprite termbox.Cell) *object {
	return &object{
		name:   name,
		index:  index,
		pos:    pos,
		sprite: sprite,
	}
}

func (o *object) ID() string {
	return fmt.Sprintf("%s_%d", o.name, o.index)
}

func (o *object) Index() int {
	return o.index
}

func (o *object) JSON() ObjectJSON {
	return ObjectJSON{
		Name:   o.name,
		Index:  o.index,
		Pos:    o.pos,
		Sprite: o.sprite,
	}
}

func (o *object) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.JSON())
}

func (o *object) Name() string {
	return o.name
}

func (o *object) Pos() *Pos {
	return o.pos
}

func (o *object) SetIndex(i int) {
	if i > 0 {
		o.index = i
	}
}

func (o *object) SetPos(p Pos) {
	o.pos = &p
}

func (o *object) Sprite() termbox.Cell {
	return o.sprite
}

func (o *object) UnmarshalJSON(data []byte) error {
	tmp := &ObjectJSON{}
	err := json.Unmarshal(data, tmp)
	if err != nil {
		return err
	}
	o.name = tmp.Name
	o.index = tmp.Index
	o.pos = tmp.Pos
	o.sprite = tmp.Sprite
	return nil
}

type ObjectJSON struct {
	Name   string
	Index  int
	Pos    *Pos
	Sprite termbox.Cell
}
