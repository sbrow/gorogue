package gorogue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestLoadSaveKeybinds(t *testing.T) {
	want := strings.Split(`[,
	{"key":"b","action":"Move","args":{"string":"SouthWest"}},
	{"key":"esc","action":"Quit"},
	{"key":"h","action":"Move","args":{"string":"West"}},
	{"key":"j","action":"Move","args":{"string":"South"}},
	{"key":"k","action":"Move","args":{"string":"North"}},
	{"key":"l","action":"Move","args":{"string":"East"}},
	{"key":"n","action":"Move","args":{"string":"SouthEast"}},
	{"key":"u","action":"Move","args":{"string":"NorthEast"}},
	{"key":"y","action":"Move","args":{"string":"NorthWest"}},
],`, "\n")
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
	got := strings.Split(string(data), "\n")
	for i, line := range got {
		if !strings.HasSuffix(line, ",") {
			got[i] += ","
		}
	}
	sort.Strings(want)
	sort.Strings(got)
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		// if !reflect.DeepEqual(got, want) {
		fmt.Printf("Wanted:\"\n%s\"\nGot:\"\n%s\"\n", strings.Join(want, "\n"), strings.Join(got, "\n"))
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
