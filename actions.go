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

func (b *BaseAction) Name() string {
	return b.name
}

func (b *BaseAction) Caller() string {
	return b.caller
}

func (b *BaseAction) Args() []interface{} {
	return b.args
}

// ActionResponse is a stuct returned from the server to a client or NPC,
// informing them of whether their action was performed, and if not, why not.
type ActionResponse struct {
	Msg   *string // Message to display to the user / logs.
	Reply bool    // Whether the action completed sucessfully
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
