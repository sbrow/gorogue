package example

import (
	"errors"
	engine "github.com/sbrow/gorogue"
	"math/rand"
)

type World struct {
	maps    []*Map
	players map[string]engine.Actor
}

func NewWorld(m ...*Map) engine.World {
	return &World{
		maps:    m,
		players: map[string]engine.Actor{},
	}
}

func (w *World) HandleAction(a *engine.Action, reply *string) error {
	var err error
	switch a.Name {
	case "Move":
		err = w.Move(a)
	case "Spawn":
		sa := &engine.Spawn{}
		sa.Caller = a.Caller
		sa.Actor = a.Args[0].(engine.Actor)
		err = w.Spawn(sa)
	}
	if err != nil {
		*reply = err.Error()
	}
	return err
}

func (w *World) Maps() []engine.Map {
	maps := []engine.Map{}
	for _, m := range w.maps {
		maps = append(maps, m)
	}
	return maps
}

func (w *World) Move(a *engine.Action) error {
	var ma engine.Move
	// TODO: Make converter from Action to *Move
	var p engine.Pos
	if a.Name != "Move" {
		//TODO: ErrorWrongAction or something.
		return errors.New("ErrorWrongAction")
	}
	caller := w.players[a.Caller]
	ma.Target = caller
	engine.Log.Println(a.Caller, caller)
	p = *caller.Pos()
	switch a.Args[0].(type) {
	case engine.Direction:
		dir := a.Args[0].(engine.Direction)
		if dir&engine.North == engine.North {
			p.Y--
		} else {
			if dir&engine.South == engine.South {
				p.Y++
			}
		}
		if dir&engine.East == engine.East {
			p.X++
		} else {
			if dir&engine.West == engine.West {
				p.X--
			}
		}
	case engine.Pos:
		p = a.Args[0].(engine.Pos)
	default:
		panic("Passed wrong args to Client.Move()")
	}
	ma.Pos = p
	caller.SetPos(&ma.Pos)
	return nil
}

func (w *World) Players() map[string]engine.Actor {
	return w.players
}

func (w *World) Spawn(s *engine.Spawn) error {
	//
	m := rand.Intn(len(w.maps))
	//
	x := rand.Intn(w.maps[m].Width)
	y := rand.Intn(w.maps[m].Height)
	p := engine.NewPos(x, y, m)
	//
	w.players[s.Caller] = s.Actor
	a := w.players[s.Caller]
	a.SetPos(p)
	w.maps[m].Players["Player_1"] = a
	engine.Log.Println(s.Caller, w.players[s.Caller])
	return nil
}
