package gorogue

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	. "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/action"
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
	quit := action.New("Quit", "Client")
	moveNorth := action.New("Move", "Client", North)
	moveNorthEast := action.New("Move", "Client", NorthEast)
	moveNorthWest := action.New("Move", "Client", NorthWest)
	moveEast := action.New("Move", "Client", East)
	moveSouth := action.New("Move", "Client", South)
	moveSouthEast := action.New("Move", "Client", SouthEast)
	moveWest := action.New("Move", "Client", West)
	moveSouthWest := action.New("Move", "Client", SouthWest)

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

// Command is a string mapped to an Action. Commands are  intended to be called
// in a Vi style command bar.
//
// TODO: (7) Implement Vi command bar.
type Command string

// Key is a keyboard key mapped to an Action. See package github.com/nsf/termbox-go
// for more information.
type Key struct {
	Mod termbox.Modifier // One of termbox.Mod* constants or 0.
	Key termbox.Key      // One of termbox.Key* constants, invalid if 'Ch' is not 0.
	Ch  rune             // a unicode character.
}

type KeyNotBoundError struct {
	K Key
}

func (k *KeyNotBoundError) Error() string {
	return fmt.Sprintf("Key %v not bound", k.K)
}
