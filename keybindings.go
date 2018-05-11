package gorogue

import (
	"errors"
	termbox "github.com/nsf/termbox-go"
)

const KeyNotBoundError string = "Key not bound"

var Commands map[Command]KeyAction

var Keybinds map[Key]KeyAction

func init() {
	// Keys
	ESC := Key{0, termbox.KeyEsc, 0}

	//KeyActions
	quit := KeyAction{"Quit", nil}
	moveNorth := KeyAction{"Move", []interface{}{North}}
	moveNorthEast := KeyAction{"Move", []interface{}{NorthEast}}
	moveNorthWest := KeyAction{"Move", []interface{}{NorthWest}}
	moveEast := KeyAction{"Move", []interface{}{East}}
	moveSouth := KeyAction{"Move", []interface{}{South}}
	moveSouthEast := KeyAction{"Move", []interface{}{SouthEast}}
	moveWest := KeyAction{"Move", []interface{}{West}}
	moveSouthWest := KeyAction{"Move", []interface{}{SouthWest}}

	Keybinds = map[Key]KeyAction{
		ESC:            quit,
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

func BindCommand(cmd Command, KeyAction KeyAction) {
	Commands[cmd] = KeyAction
}
func BindKey(key Key, KeyAction KeyAction) {
	Keybinds[key] = KeyAction
}

type KeyAction struct {
	Func string
	Args []interface{}
}

func KeyPressed(key Key) (*KeyAction, error) {
	if act, ok := Keybinds[key]; ok {
		return &act, nil
	} else {
		return nil, errors.New(KeyNotBoundError)
	}

}

func Input() (*KeyAction, error) {

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

type Command string

type Key struct {
	Mod termbox.Modifier
	Key termbox.Key
	Ch  rune
}
