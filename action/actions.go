package action

import (
	"fmt"
	engine "github.com/sbrow/gorogue"
)

func New(name, caller string, args ...interface{}) engine.Action {
	return &baseAction{name: name, caller: caller, args: args}
}

type baseAction struct {
	name   string
	caller string
	args   []interface{}
}

func (b *baseAction) Args() []interface{} {
	return b.args
}

func (b *baseAction) Caller() string {
	return b.caller
}

func (b *baseAction) Name() string {
	return b.name
}

type Quit struct {
	Caller string
}

func (q *Quit) Action() engine.Action {
	return New("Quit", q.Caller)
}

// Move  moves the target Object to the given position.
type Move struct {
	Caller string     // The originator of the Action.
	Target string     // The name of the Actor to move.
	Pos    engine.Pos // The coordinates to move them to.
}

func (m *Move) Action() engine.Action {
	return New("Move", m.Caller, m.Target, m.Pos)
}

func (m *Move) String() string {
	a := New("Move", m.Caller, m.Target, m.Pos)
	return fmt.Sprint(a)
}
