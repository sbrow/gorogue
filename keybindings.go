package gorogue

import (
	"encoding/json"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

func init() {
	fmt.Println(os.Getwd())
	fmt.Println(os.Hostname())
	if err := LoadKeyBinds(filepath.Join(basePath, "keybinds.json")); err != nil {
		panic(err)
	}
}

// Commands stores all currently bound commands.
// var Commands map[Command]Action

// Keybinds stores all currently bound keys.
var Keybinds map[Key]Action

var basePath string = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "sbrow", "gorogue")

// BindCommand maps a Command to an Action, overwriting any Action the
// Command was previously mapped to.
/*func BindCommand(cmd Command, Action Action) {
	Commands[cmd] = Action
}*/

// BindKey maps a Key to an Action, overwriting any Action the
// Key was previously mapped to.
func BindKey(key Key, Action Action) {
	Keybinds[key] = Action
}

// Input polls the user for a Key and returns the Action it's mapped to.
// If the Key pressed is bound to an Action, the action is returned along with a nil error.
// if the Key pressed isn't bound to an Action, a nil Action is returned along
// with a KeyNotBoundError.
//
// TODO: (10) Move Input() to ui package?
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

func LoadKeyBinds(path string) error {
	var tmp []KeyBind
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	Keybinds = map[Key]Action{}
	for _, bind := range tmp {
		Keybinds[bind.Key] = bind.Action
	}
	return nil
}

func SaveKeyBinds() error {
	var bind *KeyBind
	f, err := os.Create("keybinds.json")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("[\n")
	i := 0
	for k, v := range Keybinds {
		f.WriteString("\t")
		bind = &KeyBind{k, v}
		data, err := bind.MarshalJSON()
		if err != nil {
			return err
		}
		f.Write(data)
		if i+1 != len(Keybinds) {
			f.WriteString(",")
			i++
		}
		f.WriteString("\n")
	}
	f.WriteString("]")
	return nil
}

// Command is an alternate way to call an action. If a player forgets the keyboard
// shortcut for an Action, they can instead bring up the command bar with ':'
// and type in the command string for the action.
//
// TODO: (7) Implement command bar.
// type Command string

// Key represents a keyboard key or combination of keys.
//
// See package github.com/nsf/termbox-go for more information.
type Key struct {
	Mod termbox.Modifier // One of termbox.Mod* constants or 0.
	Key termbox.Key      // One of termbox.Key* constants, invalid if 'Ch' is not 0.
	Ch  rune             // a unicode character.
}

func NewKey(str string) Key {
	if len(str) == 1 {
		return Key{0, 0, rune(str[0])}
	}
	if str == "esc" {
		return Esc
	}
	return Key{}
}

func (k *Key) String() string {
	str := ""
	if k.Mod != 0 {
		// Do the thing
	}
	if k.Key != 0 {
		switch k.Key {
		case Esc.Key:
			str += "esc"
		}
	}
	if k.Ch != 0 {
		str += string(k.Ch)
	}
	return str
}

type KeyBind struct {
	Key    Key
	Action Action
}

// TODO: Try marshaling args into map[string]interface{}, see if that helps.
type KeyBindJSON struct {
	Key    string                     `json:"key"`
	Action string                     `json:"action"`
	Args   map[string]json.RawMessage `json:"args,omitempty"`
}

func (k *KeyBind) MarshalJSON() ([]byte, error) {
	var err error
	obj := &KeyBindJSON{
		Key:    k.Key.String(),
		Action: k.Action.Name,
	}
	obj.Args = map[string]json.RawMessage{}
	for _, a := range k.Action.Args {
		t := reflect.TypeOf(a)
		v, ok := a.(json.Marshaler)
		if ok {
			if obj.Args[t.Name()], err = v.MarshalJSON(); err != nil {
				return []byte{}, err
			}
		} else {
			if obj.Args[t.Name()], err = json.Marshal(a); err != nil {
				return []byte{}, err
			}
		}
	}
	return json.Marshal(obj)
}

func (k *KeyBind) UnmarshalJSON(data []byte) error {
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	k.Key = NewKey(tmp["key"].(string))
	k.Action = *NewAction(tmp["action"].(string), "Client")
	if _, ok := tmp["args"]; ok {
		args := tmp["args"].(map[string]interface{})
		for _, a := range args {
			k.Action.Args = append(k.Action.Args, a)
		}
	}
	return nil
}

// TODO: finish adding.
var (
	Backspace  Key = Key{0, termbox.KeyBackspace, 0}
	Backspace2 Key = Key{0, termbox.KeyBackspace2, 0}
	Delete     Key = Key{0, termbox.KeyDelete, 0}
	Enter      Key = Key{0, termbox.KeyEnter, 0}
	Esc        Key = Key{0, termbox.KeyEsc, 0}
	Space      Key = Key{0, termbox.KeySpace, 0}
	Tab        Key = Key{0, termbox.KeyTab, 0}
)

// KeyNotBoundError is returned when a key is looked up, but it not currently
// bound to an action.
type KeyNotBoundError struct {
	K Key
}

func (k *KeyNotBoundError) Error() string {
	return fmt.Sprintf("Key %v not bound", k.K)
}
