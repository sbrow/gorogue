package gorogue

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

// Keys
//
// TODO: finish adding.
var (
	Esc       Key = Key{0, termbox.KeyEsc, 0}
	Tab           = Key{0, termbox.KeyTab, 0}
	Space         = Key{0, termbox.KeySpace, 0}
	Backspace     = Key{0, termbox.KeyBackspace2, 0}
	Delete        = Key{0, termbox.KeyDelete, 0}
	Enter         = Key{0, termbox.KeyEnter, 0}
)

// Commands stores all currently bound commands.
var Commands map[Command]Action

// Keybinds stores all currently bound keys.
var Keybinds map[Key]Action

func init() {
	//Actions
	quit := *NewAction("Quit", "Client")
	moveNorth := *NewAction("Move", "Client", North)
	moveNorthEast := *NewAction("Move", "Client", NorthEast)
	moveNorthWest := *NewAction("Move", "Client", NorthWest)
	moveEast := *NewAction("Move", "Client", East)
	moveSouth := *NewAction("Move", "Client", South)
	moveSouthEast := *NewAction("Move", "Client", SouthEast)
	moveWest := *NewAction("Move", "Client", West)
	moveSouthWest := *NewAction("Move", "Client", SouthWest)

	Keybinds = map[Key]Action{
		Esc:            quit,
		Key{0, 0, 'k'}: moveNorth,
		Key{0, 0, 'u'}: moveNorthEast,
		Key{0, 0, 'y'}: moveNorthWest,
		Key{0, 0, 'b'}: moveSouthWest,
		Key{0, 0, 'n'}: moveSouthEast,
		Key{0, 0, 'l'}: moveEast,
		Key{0, 0, 'j'}: moveSouth,
		Key{0, 0, 'h'}: moveWest,
	}
}

// BindCommand maps a Command to an Action, overwriting any Action the
// Command was previously mapped to.
func BindCommand(cmd Command, Action Action) {
	Commands[cmd] = Action
}

// BindKey maps a Key to an Action, overwriting any Action the
// Key was previously mapped to.
func BindKey(key Key, Action Action) {
	Keybinds[key] = Action
}

// Input polls the user for a Key, and returns the Action it's mapped to or nil
// if the key isn't mapped.
func Input() (*Action, error) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			key := &Key{}
			key.Mod = ev.Mod
			if ev.Ch != 0 {
				key.Ch = ev.Ch
			} else {
				key.Key = ev.Key
			}
			return KeyPressed(*key)
		}
	}
	return nil, nil
}

// KeyPressed checks to see what action the given key is bound to.
// It returns a KeyNotBoundError if the key is unbound.
func KeyPressed(key Key) (*Action, error) {
	if act, ok := Keybinds[key]; ok {
		return &act, nil
	} else {
		return nil, &KeyNotBoundError{key}
	}

}

// Command is an alternate way to call an action. If a player forgets the keyboard
// shortcut for an Action, they can instead bring up the command bar with ':'
// and type in the command string for the action.
//
// TODO: (7) Implement command bar.
type Command string

// Key is a keyboard key mapped to an Action. See package github.com/nsf/termbox-go
// for more information.
type Key struct {
	Mod termbox.Modifier // One of termbox.Mod* constants or 0.
	Key termbox.Key      // One of termbox.Key* constants, invalid if 'Ch' is not 0.
	Ch  rune             // a unicode character.
}

// KeyNotBoundError is returned when a key is looked up, but it not currently
// bound to an action.
type KeyNotBoundError struct {
	K Key
}

func (k *KeyNotBoundError) Error() string {
	return fmt.Sprintf("Key %v not bound", k.K)
}
