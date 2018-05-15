package example

import (
	"fmt"
	engine "github.com/sbrow/gorogue"
)

// MoveAction is used to move Objects.
type MoveAction struct {
	Caller string     // The originator of the MocwAction.
	Target string     // The name of the Actor to move.
	Pos    engine.Pos // The coordinates to move them to.
}

func (m *MoveAction) String() string {
	a := engine.NewAction("Move", m.Caller, m.Target, m.Pos)
	return fmt.Sprint(a)
}

// SpawnAction is used to spawn Objects.
type SpawnAction struct {
	Caller string
	Actors Actors // TODO: (8) Change to object rather than Actor.
}
