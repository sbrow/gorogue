package gorogue

// Actions are the heart of the game's operation.
type Action interface {
	Name() string        // The name of the Action to be performed.
	Caller() string      // The originator of the Action.
	Args() []interface{} // Any relevant parameters the action needs.
}

type BaseAction struct {
	name   string
	caller string
	args   []interface{}
}

func NewAction(name, caller string, args ...interface{}) *BaseAction {
	return &BaseAction{name: name, caller: caller, args: args}
}

func (b *BaseAction) Args() []interface{} {
	return b.args
}

func (b *BaseAction) Caller() string {
	return b.caller
}

func (b *BaseAction) Name() string {
	return b.name
}

type ActionQuit struct {
	Caller string
}

func (a *ActionQuit) Action() Action {
	return &BaseAction{
		name:   "Quit",
		caller: a.Caller,
	}
}
