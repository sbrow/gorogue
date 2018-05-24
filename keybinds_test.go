package gorogue

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"testing"
)

func TestLoadSaveKeybinds(t *testing.T) {
	want := `[
	{"key":"h","action":"Move","args":{"string":"West"}},
	{"key":"k","action":"Move","args":{"string":"North"}},
	{"key":"n","action":"Move","args":{"string":"SouthEast"}},
	{"key":"b","action":"Move","args":{"string":"SouthWest"}},
	{"key":"esc","action":"Quit"},
	{"key":"y","action":"Move","args":{"string":"NorthWest"}},
	{"key":"j","action":"Move","args":{"string":"South"}},
	{"key":"u","action":"Move","args":{"string":"NorthEast"}},
	{"key":"l","action":"Move","args":{"string":"East"}}
]`
	filename := "keybinds_test.json"

	f, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	f.WriteString("[\n")
	i := 0
	var bind *KeyBind
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
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		fmt.Printf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", want, string(data))
	}
}

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
