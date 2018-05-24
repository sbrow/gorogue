package gorogue

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"testing"
)

func TestSaveKeybinds(t *testing.T) {
	var bind *KeyBind
	f, err := os.Create("keybinds.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	f.WriteString("[\n")
	i := 0
	for k, v := range Keybinds {
		f.WriteString("\t")
		bind = &KeyBind{k, v}
		data, err := bind.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		f.Write(data)
		if i+1 != len(Keybinds) {
			f.WriteString(",")
			i++
		}
		f.WriteString("\n")
	}
	f.WriteString("]")
}

// func TestLoadKeybinds(t *testing.T) {
// 	var binds []KeyBind
// 	data, err := ioutil.ReadFile("keybinds.json")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if err := json.Unmarshal(data, &binds); err != nil {
// 		t.Fatal(err)
// 	}
// 	keybinds := map[Key]Action{}
// 	for _, bind := range binds {
// 		keybinds[bind.Key] = bind.Action
// 	}
// 	fmt.Println(keybinds)
// }

func TestKeybindJSON(t *testing.T) {
	action := *NewAction("Move", "Client", South)
	key := NewKey("j")
	bind := &KeyBind{Key: key, Action: action}
	if err := JSONTester(bind, &KeyBind{}); err != nil {
		t.Fatal(err)
	}
}

func TestDirectionJSON(t *testing.T) {
	dir := South
	var obj Direction
	err := JSONTester(dir, &obj)
	if err != nil {
		t.Fatal(err)
	}
}

func JSONTester(obj interface{}, out interface{}) error {
	fmt.Println("pre ", obj)

	byt, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Println("byte", string(byt))

	if err = json.Unmarshal(byt, &out); err != nil {
		return err
	}
	fmt.Println("Post", out)
	return nil
}
