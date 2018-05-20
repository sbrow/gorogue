package gorogue

import (
	"errors"
	"math/rand"
)

type World struct {
	Maps    []*Map
	Players map[string]Actor
}

func NewWorld(m ...*Map) *World {
	return &World{
		Maps:    m,
		Players: map[string]Actor{},
	}
}

func (w *World) HandleAction(a *Action, reply *string) error {
	var err error
	switch a.Name {
	case "Move":
		err = w.Move(a)
	case "Spawn":
		sa := &SpawnAction{}
		sa.Caller = a.Caller
		sa.Actor = a.Args[0].(Actor)
		err = w.Spawn(sa)
	}
	if err != nil {
		msg := err.Error()
		reply = &msg
	}
	return err
}

/*func (w *World) Maps() []*Map {
	return w.Maps
}*/

func (w *World) Move(a *Action) error {
	var ma MoveAction
	var p Pos
	if a.Name != "Move" {
		//TODO: ErrorWrongAction or something.
		return errors.New("ErrorWrongAction")
	}
	caller := w.Players[a.Caller]
	ma.Target = caller
	p = *caller.Pos()
	switch a.Args[0].(type) {
	case Direction:
		pt := a.Args[0].(Direction).Point()
		p.X += pt.X
		p.Y += pt.Y
	case Pos:
		p = a.Args[0].(Pos)
	default:
		return errors.New("Passed wrong args to Client.Move()")
	}
	ma.Pos = p
	Log.Println(a)
	w.Maps[ma.Pos.Map].Move(ma)
	return nil
}

/*func (w *World) Players() map[string]Actor {
	return w.Players
}*/

func (w *World) Spawn(s *SpawnAction) error {
	//
	m := rand.Intn(len(w.Maps))
	//
	x := rand.Intn(w.Maps[m].Width)
	y := rand.Intn(w.Maps[m].Height)
	p := NewPos(x, y, m)
	//
	w.Players[s.Caller] = s.Actor
	a := w.Players[s.Caller]
	a.SetPos(p)
	w.Maps[m].Players["Player_1"] = a
	Log.Println(s.Caller, w.Players[s.Caller])
	return nil
}
