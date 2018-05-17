package gorogue_test

import (
	"encoding/json"
	"fmt"
	// termbox "github.com/nsf/termbox-go"
	// "testing"
)

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
	fmt.Println("post", out)
	return nil
}

/*
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

func TestActions(t *testing.T) {
	pos := Pos{Point{3, 5}, "Map"}
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
*/
/*
func TestMoreJSON(t *testing.T) {
	n := NewPlayer("Player", 1)
*/
/*	m := map[string]*player{
		"Player_1": n.(*player),
	}
*/
/*
m := map[string]Player{
		"Player_1": n.(*player),
	}
	data, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		t.Fatal(err)
	}
	fmt.Println(tmp)
	fmt.Println(Mapper(tmp)["Player_1"])
}

func Mapper(v map[string]interface{}) map[string]Player {
	out := map[string]Player{}
	for j := range v {
		p := v[j].(map[string]interface{})
		tmp := &player{}
		for k, val := range p {
			if val != nil {
				switch k {
				case "Name":
					tmp.name = val.(string)
				case "Index":
					tmp.index = int(val.(float64))
				case "Pos":
					tmp.pos = AsPos(val.(map[string]interface{}))
				case "HP":
					tmp.hp = int(val.(float64))
				}
			}
		}
		out[j] = tmp
	}
	return out
}

func AsPos(v map[string]interface{}) *Pos {
	for k, v := range v {
		fmt.Println("thing", k, v)
	}
	return nil
}
*/

// func TestBorder(t *testing.T) {
// 	termbox.Init()
// 	defer termbox.Close()

// 	// Top-Left corner.
// 	Ox, Oy := 0, 0
// 	// Bottom-Right corner.
// 	w, h := 19, 11

// 	s := []rune(LightBorder)

// 	// Print the horizontals
// 	for x := 1; x < w-1; x++ {
// 		termbox.SetCell(x, Oy, s[0], termbox.ColorDefault, termbox.ColorBlack)
// 		termbox.SetCell(x, h/2, s[0], termbox.ColorDefault, termbox.ColorBlack)
// 		termbox.SetCell(x, h-1, s[0], termbox.ColorDefault, termbox.ColorBlack)
// 	}
// 	// Print the verticals
// 	for y := 1; y < h-1; y++ {
// 		termbox.SetCell(Ox, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
// 		termbox.SetCell(w/2, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
// 		termbox.SetCell(w-1, y, s[1], termbox.ColorDefault, termbox.ColorBlack)
// 	}

// 	// Print the corners.
// 	termbox.SetCell(Ox, Oy, s[2], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w/2, Oy, s[3], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w-1, Oy, s[4], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(Ox, h/2, s[5], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w/2, h/2, s[6], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w-1, h/2, s[7], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(Ox, h-1, s[8], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w/2, h-1, s[9], termbox.ColorDefault, termbox.ColorBlack)
// 	termbox.SetCell(w-1, h-1, s[10], termbox.ColorDefault, termbox.ColorBlack)

// 	termbox.Flush()
// 	_ = termbox.PollEvent()
// }
