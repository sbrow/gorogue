package gorogue

import (
	"encoding/json"
	termbox "github.com/nsf/termbox-go"
)

type Object interface {
	Name() string
	MarshalJSON() ([]byte, error)
	Pos() *Point
	SetPos(pt Point)
	Sprite() termbox.Cell
	UnmarshalJSON(data []byte) error
}

type object struct {
	name   string
	pos    *Point
	sprite termbox.Cell
}

func newObject(name string, pos *Point, sprite termbox.Cell) *object {
	return &object{
		name:   name,
		pos:    pos,
		sprite: sprite,
	}
}

func (o *object) JSON() ObjectJSON {
	return ObjectJSON{
		Name:   o.name,
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

func (o *object) Pos() *Point {
	return o.pos
}

func (o *object) SetPos(pt Point) {
	o.pos = &pt
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
	o.pos = tmp.Pos
	o.sprite = tmp.Sprite
	return nil
}

type ObjectJSON struct {
	Name   string
	Pos    *Point
	Sprite termbox.Cell
}
