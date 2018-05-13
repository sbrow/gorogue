package example

import rogue "github.com/sbrow/gorogue"

// MoveAction is used to move Objects.
type MoveAction struct {
	Caller string    // The originator of the MocwAction.
	Target string    // The name of the Actor to move.
	Pos    rogue.Pos // The coordinates to move them to.
}

// MoveAction is used to spawn Objects.
type SpawnAction struct {
	Caller string
	Actors Actors // TODO: (8) Change to object rather than Actor.
}
