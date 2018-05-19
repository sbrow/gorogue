package gorogue

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	name   string
	index  int
	pos    *Pos
	sprite Sprite
}

// NewPlayer creates a new Player using the standard '@' character sprite.
func NewPlayer(name string) *Player {
	return &Player{
		name:   name,
		index:  1,
		pos:    nil,
		sprite: DefaultPlayer,
	}
}

func (p *Player) ID() string {
	return fmt.Sprintf("%s_%d", p.name, p.index)
}

func (p *Player) Index() int {
	return p.index
}

func (p *Player) JSON() PlayerJSON {
	return PlayerJSON{
		Name:   p.name,
		Index:  p.index,
		Pos:    p.pos,
		Sprite: p.sprite,
		Type:   "Player",
	}
}

// MarshalJSON converts the Player into JSON bytes.
func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.JSON())
}

func (p *Player) Move(pos Pos) bool {
	p.SetPos(&pos)
	return true
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Pos() *Pos {
	return p.pos
}

func (p *Player) SetIndex(i int) {
	if i > 0 {
		p.index = i
	}
}

func (p *Player) SetPos(pos *Pos) {
	p.pos = pos
}

func (p *Player) Sprite() Sprite {
	return p.sprite
}

// UnmarshalJSON reads JSON data into this Player.
func (p *Player) UnmarshalJSON(data []byte) error {
	tmp := &PlayerJSON{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	p.name = tmp.Name
	p.index = tmp.Index
	p.pos = tmp.Pos
	p.sprite = tmp.Sprite
	return nil
}

// PlayerJSON allows Player objects to be converted into JSON
// and transported via JSON RPC.
type PlayerJSON struct {
	Name   string
	Index  int
	Pos    *Pos
	Sprite Sprite
	Type   string
}
