package assets

import (
	// "encoding/json"
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/sbrow/gorogue"
)

type Player struct {
	name   string
	id     int
	mp     *Map
	ready  bool
	init   int
	pt     *Point
	ch     chan string
	sprite termbox.Cell
}

// NewPlayer creates a new Player using the standard '@' character sprite.
func NewPlayer(name string) *Player {
	p := &Player{}
	p.name = name
	p.id = 1
	p.pt = nil
	p.sprite = DefaultPlayer
	p.ch = make(chan string)

	return p
}

func (p *Player) Done() {
	p.ch <- "Done"
	p.ready = false
}

func (p *Player) ID() string {
	return fmt.Sprintf("%s_%d", p.name, p.id)
}

func (p *Player) Index() int {
	return p.id
}

/*func (p *Player) JSON() PlayerJSON {
	return PlayerJSON{
		Name:   p.name,
		Index:  p.id,
		Pos:    p.pt,
		Sprite: p.sprite,
		Type:   "Player",
	}
}
*/
// MarshalJSON converts the Player into JSON bytes.
/*func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.JSON())
}
*/

func (p *Player) Map() *Map {
	return p.mp
}

func (p *Player) Move(pos *Pos) error {
	p.SetPos(pos)
	return nil
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Pos() *Pos {
	return &Pos{*p.pt, p.mp.Index}
}

func (p *Player) Ready() {
	if p.ready {
		return
	}
	p.ready = true
	p.mp.Ready <- NewItem(p.ch, p.init)
	if _, ok := <-p.ch; !ok {
		panic("Something wrong with the channel.")
	}
}

func (p *Player) SetIndex(i int) {
	if i > 0 {
		p.id = i
	}
}

func (p *Player) SetMap(m *Map) {
	p.mp = m
}

func (p *Player) SetPos(pos *Pos) error {
	if p.mp == nil {
		return errors.New("Player does not have a map!")
	}
	p.mp = p.mp.World.Maps()[pos.Map]
	p.pt = &pos.Point
	return nil
}

func (p *Player) SetPt(pt *Point) {
	p.pt = pt
}

func (p *Player) Sprite() termbox.Cell {
	return p.sprite
}

// UnmarshalJSON reads JSON data into this Player.
/*func (p *Player) UnmarshalJSON(data []byte) error {
	tmp := &PlayerJSON{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	p.name = tmp.Name
	p.id = tmp.Index
	p.pt = tmp.Pos
	p.sprite = tmp.Sprite
	return nil
}*/

// PlayerJSON allows Player objects to be converted into JSON
// and transported via JSON RPC.
type PlayerJSON struct {
	Name   string
	Index  int
	Pos    *Pos
	Sprite termbox.Cell
	Type   string
}
