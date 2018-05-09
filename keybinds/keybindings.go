package keybinds

import (
	"errors"
	termbox "github.com/nsf/termbox-go"
	rogue "github.com/sbrow/gorogue"
)

const KeyNotBoundError string = "Key not bound"

var Commands map[Command]Action

var Keybinds map[Key]Action

func init() {
	// Keys
	ESC := Key{0, termbox.KeyEsc, 0}

	//Actions
	quit := Action{"Quit", nil}
	moveNorth := Action{"Move", []interface{}{rogue.North}}
	moveNorthEast := Action{"Move", []interface{}{rogue.NorthEast}}
	moveEast := Action{"Move", []interface{}{rogue.East}}
	moveSouth := Action{"Move", []interface{}{rogue.South}}
	moveWest := Action{"Move", []interface{}{rogue.West}}

	Keybinds = map[Key]Action{
		ESC:            quit,
		Key{0, 0, 'k'}: moveNorth,
		Key{0, 0, 'u'}: moveNorthEast,
		Key{0, 0, 'l'}: moveEast,
		Key{0, 0, 'j'}: moveSouth,
		Key{0, 0, 'h'}: moveWest,
	}
}

func BindCommand(cmd Command, action Action) {
	Commands[cmd] = action
}
func BindKey(key Key, action Action) {
	Keybinds[key] = action
}

type Action struct {
	Func string
	Args []interface{}
}

func KeyPressed(key Key) (*Action, error) {
	if act, ok := Keybinds[key]; ok {
		return &act, nil
	} else {
		return nil, errors.New(KeyNotBoundError)
	}

}

func Input() (*Action, error) {

	for {
		// TODO: Pass event to Keybindings
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

type Command string

type Key struct {
	Mod termbox.Modifier
	Key termbox.Key
	Ch  rune
}
