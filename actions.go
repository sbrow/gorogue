package gorogue

type Action struct {
	Name   string
	Caller string
	Args   []interface{}
}

func NewAction(name, caller string, args ...interface{}) *Action {
	return &Action{Name: name, Caller: caller, Args: args}
}

type Quit struct {
	Caller string
}

func (q *Quit) Action() *Action {
	return NewAction("Quit", q.Caller)
}

// Move  moves the target Object to the given position.
type Move struct {
	Caller string // The originator of the Action.
	Target Object // The name of the Actor to move.
	Pos    Pos    // The coordinates to move them to.
}

func (m *Move) Action() *Action {
	return NewAction("Move", m.Caller, m.Target, m.Pos)
}

// Spawn is used to spawn Objects.
type Spawn struct {
	Caller string
	Actor  Actor // TODO: (8) Change to object rather than Actor.
}
