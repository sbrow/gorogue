package gorogue

import (
	"encoding/json"
	"github.com/sbrow/gorogue/sprites"
)

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
			actual = &player{}
		}

		err = json.Unmarshal(r, actual)
		if err != nil {
			return err
		}
		*a = append(*a, actual)

	}
	return nil
}

type NPC struct {
	object
	hp int
}

// Player is any playable character. Players are controlled by clients.
// Each client can control more than one Player.
type Player interface {
	Actor
	HP() int
}

type player struct {
	object
	hp int
}

// NewPlayer creates a new player using the standard '@' character sprite.
func NewPlayer(name string, hp int) Player {
	return &player{
		object: *newObject(name, 1, nil, sprites.Default),
		hp:     hp,
	}
}

func (p *player) JSON() PlayerJSON {
	return PlayerJSON{
		Type:       "player",
		ObjectJSON: p.object.JSON(),
		HP:         p.hp,
	}
}

func (p *player) HP() int {
	return p.hp
}

// MarshalJSON converts the Player into JSON bytes.
func (p *player) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.JSON())
}

func (p *player) Move(pos Pos) bool {
	p.SetPos(pos)
	return true
}

// UnmarshalJSON reads JSON data into this Player.
func (p *player) UnmarshalJSON(data []byte) error {
	tmp := &PlayerJSON{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	byt, err := json.Marshal(tmp.ObjectJSON)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byt, &p.object)
	if err != nil {
		return err
	}
	p.hp = tmp.HP
	return nil
}

// PlayerJSON allows Player objects to be converted into JSON
// and transported via JSON RPC.
type PlayerJSON struct {
	ObjectJSON
	Type string
	HP   int
}
