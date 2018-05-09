package keybinds

import (
	"errors"
	termbox "github.com/nsf/termbox-go"
)

const KeyNotBoundError string = "Key not bound"

var Commands map[Command]Action

var Keybinds map[Key]Action

func init() {
	// Keys
	ESC := Key{0, termbox.KeyEsc, 0}
	j := Key{0, 0, 'j'}

	//Actions
	quit := Action{"Quit", nil}
	moveEast := Action{"Move", []interface{}{"East"}} //TODO: Fix

	Keybinds = map[Key]Action{
		ESC: quit,
		j:   moveEast,
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
