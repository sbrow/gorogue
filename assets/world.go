package assets

import (
	"errors"
	"fmt"
	. "github.com/sbrow/gorogue"
	"math/rand"
)

type ExampleWorld struct {
	maps    []*Map
	Players map[string]Actor
}

func NewWorld() *ExampleWorld {
	w := &ExampleWorld{}
	w.maps = []*Map{}
	w.Players = map[string]Actor{}
	return w
}

func (e *ExampleWorld) NewMap(w, h int) {
	m := NewMap(5, 5)
	m.World = e
	e.maps = append(e.maps, m)
}

func (w *ExampleWorld) HandleAction(a *Action, reply *string) error {
	var err error
	switch a.Name {
	case "Move":
		err = w.Move(a)
	case "Spawn":
		err = w.Spawn(a)
	}
	if err != nil {
		msg := err.Error()
		reply = &msg
	}
	return err
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
	caller := w.Players[a.Caller]
	ma.Target = caller
	p := caller.Pos()
	switch a.Args[0].(type) {
	case Direction:
		pt := a.Args[0].(Direction).Point()
		p.X += pt.X
		p.Y += pt.Y
	case Pos:
		p = a.Args[0].(*Pos)
	default:
		return errors.New("Passed wrong args to Client.Move()")
	}
	ma.Pos = *p
	Log.Println(a)
	w.maps[ma.Pos.Map].Move(ma)
	return nil
}

func (w *ExampleWorld) Spawn(a *Action) error {
	sa := &SpawnAction{}
	sa.Caller = a.Caller
	sa.Actor = a.Args[0].(Actor)

	m := rand.Intn(len(w.maps))
	x := rand.Intn(w.maps[m].Width)
	y := rand.Intn(w.maps[m].Height)
	fmt.Println(x, y, m)
	p := NewPos(x, y, m)

	w.Players[sa.Caller] = sa.Actor
	actor := w.Players[sa.Caller]
	Map := w.maps[m]
	Log.Println(w.maps[m].World)
	Map.Players["Player_1"] = actor
	actor.SetMap(Map)
	actor.SetPos(*p)
	if len(Map.Players) == 1 {
		go Map.Tick()
	}
	Log.Println(sa.Caller, w.Players[sa.Caller])
	return nil
}
