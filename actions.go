package gorogue

// Action requests that the World perform some function.
// Actions are passed from Actors/Clients to Servers/Worlds,
// where they are evaluated. Almost every event that occurs
// in the game should be the result of an action.
//
// One Action is collected from each Actor every Map Tick.
type Action struct {
	Name   string        // The type of Action to perform.
	Caller string        // The name of the object that requested this Action.
	Args   []interface{} // Any other relevant parameters.
}

// NewAction returns a new, generic action with the given traits.
// Custom actions should have a function to convert them to generic
// actions, as generic actions are far less likely to have problems
// unmarshaling when transferring to a Server.
func NewAction(name, caller string, args ...interface{}) *Action {
	return &Action{Name: name, Caller: caller, Args: args}
}

// MoveAction moves the target Object to the given position.
type MoveAction struct {
	Caller string `json:"caller"`           // The originator of the Action.
	Target Object `json:"target,omitempty"` // The name of the Object to move.
	Pos    Pos    // The coordinates the Object should be moved to.
}

// Action returns a generic version of the MoveAction.
// This should be called when sending the Action to a server
// via HandleAction.
func (m *MoveAction) Action() *Action {
	return NewAction("Move", m.Caller, m.Target, m.Pos)
}

// QuitAction stops the current Client, disconnecting from the server if the
// Client is a RemoteClient.
type QuitAction struct {
	Caller string
}

// Action returns a generic version of the QuitAction.
// This should be called when sending the Action to a server
// via HandleAction.
func (q *QuitAction) Action() *Action {
	return NewAction("Quit", q.Caller)
}

// SpawnAction is used to spawn Objects.
type SpawnAction struct {
	Caller string
	Actor  Actor // TODO: (8) Change to object rather than Actor.
}
