package gorogue

// Actions are the heart of the game's operation.
type Action struct {
	Name   string        // The name of the Action to be performed.
	Caller string        // The originator of the Action.
	Args   []interface{} // Any relevant parameters the action needs.
}

// ActionResponse is a stuct returned from the server to a client or NPC,
// informing them of whether their action was performed, and if not, why not.
type ActionResponse struct {
	Msg   *string // Message to display to the user / logs.
	Reply bool    // Whether the action completed sucessfully
}

// MoveAction is used to move Objects.
type MoveAction struct {
	Caller string // The originator of the MocwAction.
	Target string // The name of the Actor to move.
	Pos    Pos    // The coordinates to move them to.
}

// MoveAction is used to spawn Objects.
type SpawnAction struct {
	Caller string
	Actors Actors // TODO: (8) Change to object rather than Actor.
}
