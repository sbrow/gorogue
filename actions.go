package gorogue

type Action struct {
	Name   string
	Caller string
	Args   []interface{}
}

type ActionResponse struct {
	Msg   string
	Reply bool
}

type MoveAction struct {
	Caller string
	Pos    Pos
}

type SpawnAction struct {
	Caller string
	Actors Actors
}
