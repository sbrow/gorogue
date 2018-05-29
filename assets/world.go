package assets

import (
	"errors"
	"math/rand"

	. "github.com/sbrow/gorogue"
)

type ExampleWorld struct {
	maps    []*Map
	players map[string]Actor
}

func NewWorld() *ExampleWorld {
	w := &ExampleWorld{}
	w.maps = []*Map{}
	w.players = map[string]Actor{}
	return w
}

func (e *ExampleWorld) NewMap(w, h int) {
	m := NewMap(w, h)
	m.World = e
	e.maps = append(e.maps, m)
}

func (e *ExampleWorld) HandleAction(a *Action, reply *string) error {
	Log.Debugln("Players:")
	for k, v := range e.players {
		Log.Printf("[%s]=%v\n", k, v)
	}
	Log.Debugln()
	var err error
	switch a.Name {
	case "Move":
		err = e.Move(a)
	case "Spawn":
		err = e.Spawn(a)
	}
	if err != nil {
		msg := err.Error()
		reply = &msg
	}
	if reply == nil {
		return nil
	}
	return errors.New(*reply)
}

func (w *ExampleWorld) Maps() []*Map {
	return w.maps
}

func (w *ExampleWorld) Move(a *Action) error {
	var ma MoveAction
	if a.Name != "Move" {
		//TODO: ErrorWrongAction or something.
		return errors.New("ErrorWrongAction")
	}
	caller := w.players[a.Caller]
	ma.Target = caller
	var pt Point
	p := caller.Pos()
	switch a.Args[0].(type) {
	case Direction:
		pt = a.Args[0].(Direction).Point()
	case Pos:
		p = a.Args[0].(*Pos)
	case string:
		pt = Dir(a.Args[0].(string)).Point()
	case float64:
		pt = Direction(uint8(a.Args[0].(float64))).Point()
	default:
		return errors.New("Passed wrong args to Client.Move()")
	}
	p.X += pt.X
	p.Y += pt.Y
	ma.Pos = *p
	w.maps[ma.Pos.Map].Move(ma)
	return nil
}

func (e *ExampleWorld) Players() map[string]Actor {
	return e.players
}
func (w *ExampleWorld) Spawn(a *Action) error {
	sa := &SpawnAction{}
	sa.Caller = a.Caller
	sa.Actor = a.Args[0].(Actor)

	m := rand.Intn(len(w.maps))
	w.players[sa.Caller] = sa.Actor

	return w.maps[m].Spawn(sa)
}
