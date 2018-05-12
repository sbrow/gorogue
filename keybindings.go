package gorogue

import (
	"errors"
	termbox "github.com/nsf/termbox-go"
)

const KeyNotBoundError string = "Key not bound"

// Commands stores all currently bound commands
var Commands map[Command]Action

// Keybinds stores all currently bound keys.
var Keybinds map[Key]Action

var (
	Esc       Key = Key{0, termbox.KeyEsc, 0}
	Tab       Key = Key{0, termbox.KeyTab, 0}
	Space     Key = Key{0, termbox.KeySpace, 0}
	Backspace Key = Key{0, termbox.KeyBackspace, 0}
	Enter     Key = Key{0, termbox.KeyEnter, 0}
)

func init() {
	//Actions
	quit := Action{Name: "Quit"}
	moveNorth := Action{Name: "Move", Caller: "Client", Args: []interface{}{North}}
	moveNorthEast := Action{Name: "Move", Caller: "Client", Args: []interface{}{NorthEast}}
	moveNorthWest := Action{Name: "Move", Caller: "Client", Args: []interface{}{NorthWest}}
	moveEast := Action{Name: "Move", Caller: "Client", Args: []interface{}{East}}
	moveSouth := Action{Name: "Move", Caller: "Client", Args: []interface{}{South}}
	moveSouthEast := Action{Name: "Move", Caller: "Client", Args: []interface{}{SouthEast}}
	moveWest := Action{Name: "Move", Caller: "Client", Args: []interface{}{West}}
	moveSouthWest := Action{Name: "Move", Caller: "Client", Args: []interface{}{SouthWest}}

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

// BindCommand maps a Command to an action, overwriting any action the
// Command was previously mapped to
func BindCommand(cmd Command, Action Action) {
	Commands[cmd] = Action
}

// BindKey maps a key to an action, overwriting any action the
// Key was previously mapped to
func BindKey(key Key, Action Action) {
	Keybinds[key] = Action
}

// Input polls the user for a Key, and returns the Action it's mapped to (if any).
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
	return nil, errors.New("Something  went wrong.")
}

func KeyPressed(key Key) (*Action, error) {
	if act, ok := Keybinds[key]; ok {
		return &act, nil
	} else {
		return nil, errors.New(KeyNotBoundError)
	}

}

// Command is a string mapped to an Action. Commands are  intended to be called
// in a Vi style command bar.
//
// TODO: (7) Implement Vi command bar.
type Command string

// Keyd is a keyboard key mapped to an Action. See package github.com/nsf/termbox-go
// for more information.
type Key struct {
	Mod termbox.Modifier // One of termbox.Mod* constants or 0.
	Key termbox.Key      // One of termbox.Key* constants, invalid if 'Ch' is not 0.
	Ch  rune             // a unicode character.
}
