package gorogue

import (
	"encoding/json"
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

var PlayerSprite termbox.Cell = termbox.Cell{'@', termbox.ColorWhite, termbox.ColorDefault}

// Actor is an object that can act freely. There are two main kinds of actors:
// player characters and non-player characters (NPCs). The important
// distinction being that NPCs are controlled by the server, and Player
// characters are controlled by clients.
//
// Each NPC gets their own goroutine, meaning each acts on their own thread,
// separate from other actors. The server receives requests to act from each NPC,
// and determines whether that action is valid. If if isn't, the action is rejected
// and the Actor must choose a different action to perform. If the action is valid,
// it gets stored in memory and is called during the next Map tick. (See Map.Tick)

type Actor interface {
	Object             // The Object interface.
	Move(pos Pos) bool // Moves the Actor to the given position.
}

// Actors is a wrapper for an array of Actors. It is necessary to
// Unmarshal objects that implement Actor.
type Actors []Actor

type NPC struct {
	object
	hp int
}

// Takes JSON data and reads it into this array.
func (a *Actors) UnmarshalJSON(data []byte) error {
	// this just splits up the JSON array into the raw JSON for each object
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, r := range raw {
		// unamrshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return err
		}

		Type := ""
		if t, ok := obj["Type"].(string); ok {
			Type = t
		}

		// unmarshal again into the correct type
		var actual Actor
		switch Type {
		case "Player":
			actual = &Player{}
		}

		err = json.Unmarshal(r, actual)
		if err != nil {
			return err
		}
		*a = append(*a, actual)

	}
	return nil
}

type Player struct {
	name   string
	index  int
	pos    *Pos
	sprite termbox.Cell
}

// NewPlayer creates a new Player using the standard '@' character sprite.
func NewPlayer(name string) *Player {
	return &Player{
		name:   name,
		index:  1,
		pos:    nil,
		sprite: PlayerSprite,
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
	p.SetPos(pos)
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

func (p *Player) SetPos(pos Pos) {
	p.pos = &pos
}

func (p *Player) Sprite() termbox.Cell {
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
	Sprite termbox.Cell
	Type   string
}
