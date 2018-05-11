package gorogue

import (
	"encoding/json"
	"fmt"
	// termbox "github.com/nsf/termbox-go"
	"testing"
)

/*
func JSONTester(obj interface{}, out interface{}) error {
	fmt.Println("pre ", obj)

	byt, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Println("byte", string(byt))

	n := out
	err = json.Unmarshal(byt, &n)
	if err != nil {
		return err
	}
	fmt.Println("post", string(byt))
	fmt.Println()
	return nil
}
func TestPlayerJSON(t *testing.T) {
	err := JSONTester(NewPlayer("PlayerOne", 1), &Player{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestActorsJSON(t *testing.T) {
	err := JSONTester(Actors([]Actor{NewPlayer("PlayerOne", 1)}), Actors([]Actor{}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNumbers(t *testing.T) {
	args := &MoveArgs{Actors([]Actor{}), []Point{Point{0, 0}}}
	dir := North
	dir = NorthWest
	fmt.Println(dir, South, dir&South == South)
	fmt.Println(dir, North, dir&North, dir&North == North)

	if dir&North == North {
		args.Points[0].Y--
	} else {
		if dir&South == South {
			args.Points[0].Y++
		}
	}
	if dir&East == East {
		args.Points[0].X++
	} else {
		if dir&West == West {
			args.Points[0].X--
		}
	}
	fmt.Println(args.Points[0])
}

func TestObjectSprite(t *testing.T) {
	p := NewPlayer("PlayerOne", 1)
	s := p.Sprite()
	fmt.Println(string(s.Ch))
	fmt.Println(p.Sprite())
}
*/
func TestActions(t *testing.T) {
	pos := Pos{Point{3, 5}, 0}
	act := Action{
		Name:   "Move",
		Caller: "Player",
		Args:   []interface{}{pos},
	}
	data, err := json.Marshal(act)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(act)
	fmt.Println(string(data))

	var a Action
	if err := json.Unmarshal(data, &a); err != nil {
		t.Fatal(err)
	}
	switch a.Name {
	case "Move":
		var ma MoveAction
		ma.Caller = a.Caller
		tmp := a.Args[0].(map[string]interface{})
		fmt.Println(tmp)
		ma.Pos = Pos{Point{int(tmp["X"].(float64)), int(tmp["Y"].(float64))}, int(tmp["Map"].(float64))}
		fmt.Println(a)
		fmt.Println(ma)
	}
}

func TestConvertPos(t *testing.T) {
	ps := Pos{Point{3, 5}, 0}

}
