package gorogue_test

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"testing"
)

func TestMoveAction(t *testing.T) {
	name := "Move"
	caller := "Testing"
	pos := Pos{Point{3, 5}, 0}
	target := NewPlayer("Player")
	action := NewAction(name, caller, target, pos)
	moveAction := MoveAction{
		Caller: caller,
		Target: target,
		Pos:    pos,
	}
	if fmt.Sprint(*action) != fmt.Sprint(*moveAction.Action()) {
		t.Fatal("MoveAction does not match.")
	}
}

func TestQuitAction(t *testing.T) {
	name := "Quit"
	caller := "Testing"
	action := NewAction(name, caller)
	quitAction := QuitAction{
		Caller: caller,
	}
	if fmt.Sprint(*action) != fmt.Sprint(quitAction.Action()) {
		t.Fatal("QuitAction does not match.")
	}
}
