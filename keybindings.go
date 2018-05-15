package gorogue

import (
	"errors"
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
	Backspace     = Key{0, termbox.KeyBackspace, 0}
	Enter         = Key{0, termbox.KeyEnter, 0}
)

// Commands stores all currently bound commands.
var Commands map[Command]Action

// Keybinds stores all currently bound keys.
var Keybinds map[Key]Action

func init() {
	//Actions
	quit := NewAction("Quit", "Client")
	moveNorth := NewAction("Move", "Client", North)
	moveNorthEast := NewAction("Move", "Client", NorthEast)
	moveNorthWest := NewAction("Move", "Client", NorthWest)
	moveEast := NewAction("Move", "Client", East)
	moveSouth := NewAction("Move", "Client", South)
	moveSouthEast := NewAction("Move", "Client", SouthEast)
	moveWest := NewAction("Move", "Client", West)
	moveSouthWest := NewAction("Move", "Client", SouthWest)

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

type KeyNotBoundError struct {
	K Key
}

func (k *KeyNotBoundError) Error() string {
	return fmt.Sprintf("Key %v not bound", k.K)
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

// Input polls the user for a Key, and returns the Action it's mapped to (if any).
func Input() (Action, error) {
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
	return nil, errors.New("Something  went wrong.")
}

func KeyPressed(key Key) (Action, error) {
	if act, ok := Keybinds[key]; ok {
		return act, nil
	} else {
		return nil, &KeyNotBoundError{key}
	}

}
