package gorogue

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMoveArgs(t *testing.T) {
	m := &MoveArgs{
		Actors: Actors([]Actor{NewPlayer("game", DefaultSprite, 1)}),
		Points: []Point{Point{1, 1}},
	}
	fmt.Println("pre Marshal", m)

	byt, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("byte", string(byt))

	var n *MoveArgs = &MoveArgs{}
	err = json.Unmarshal(byt, &n)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("unmarshalled", string(byt))
}
