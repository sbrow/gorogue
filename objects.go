package gorogue

import (
	"encoding/json"
	// "fmt"
	termbox "github.com/nsf/termbox-go"
)

// Actor is an object that can act freely. There are two main kinds of actors:
// player characters and non-player characters (NPCs). The most important
// difference between them is that NPCs are handled server-side, and Player
// characters are handled client-side.
//
// Each Actor gets their own goroutine, meaning each acts on their own thread,
// separate of all others. The server receives requests to act from each actor,
// and determines whether that action is valid. If not, the action is rejected
// and the Actor must choose a different action to take. If the action is valid,
// it gets stored in memory and is called during the next Map tick. (See Map.Tick)
//
type Actor interface {
	MarshalJSON() ([]byte, error)
	Pos() *Point
	SetPos(pt Point)
	Sprite() termbox.Cell
	UnmarshalJSON(b []byte) error
}

type Actors []Actor

func (a *Actors) UnmarshalJSON(data []byte) error {
	// this just splits up the JSON array into the raw JSON for each object
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		// s := string(data)
		// s = s[10 : len(s)-26]
		// fmt.Println("data", s)
		// err = json.Unmarshal([]byte(s), &raw)

		// var m map[string]*json.RawMessage
		// err = json.Unmarshal(data, &m)
		// fmt.Println(m)
		// fmt.Println(s)
		// fmt.Println("raw", raw)
		// if err != nil {
		// panic(err)
		// }
		// return nil
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
	pos    *Point
	sprite termbox.Cell
	hp     int
}

type PlayerJSON struct {
	Name   string
	Type   string
	Pos    *Point
	Sprite termbox.Cell
	HP     int
}

func NewPlayer(name string, Sprite termbox.Cell, hp int) *Player {
	return &Player{
		name:   "Player",
		pos:    nil,
		sprite: termbox.Cell{'@', termbox.ColorWhite, termbox.ColorDefault},
		hp:     hp,
	}
}

func (p *Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(&PlayerJSON{Type: "Player", Name: p.name,
		Pos: p.pos, Sprite: p.sprite, HP: p.hp})
}

func (p *Player) UnmarshalJSON(b []byte) error {
	tmp := &PlayerJSON{}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	p.name = tmp.Name
	p.pos = tmp.Pos
	p.hp = tmp.HP
	p.sprite = tmp.Sprite
	return nil

}

func (p *Player) Pos() *Point {
	return p.pos
}

func (p *Player) SetPos(pt Point) {
	p.pos = &pt
}

func (p *Player) Sprite() termbox.Cell {
	return p.sprite
}
