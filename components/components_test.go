package components

import (
	"testing"

	"engo.io/ecs"
	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue"
)

func TestActionFace(t *testing.T) {
	e := struct {
		ecs.BasicEntity
		Action
	}{
		ecs.NewBasic(),
		Action{Name: "Test", Caller: "Client"},
	}

	var v ActionFace
	v = &e
	if &e.Action != v.GetAction() {
		t.Fatal("Action does not implement ActionFace.")
	}
}

func TestBasicFace(t *testing.T) {
	var b ecs.BasicEntity
	b = ecs.NewBasic()
	var v BasicFace
	v = &b
	if &b != v.GetBasicEntity() {
		t.Fatal("ecs.BasicEntity does not implement BasicFace.")
	}
}

func TestPosFace(t *testing.T) {
	p := &Pos{X: 3, Y: 5}
	var v PosFace
	v = p
	if p != v.GetPos() {
		t.Fatal("Pos does not implement PosFace.")
	}
}

func TestPosString(t *testing.T) {
	p := &Pos{X: 3, Y: 5}
	want := "{3, 5}"
	if p.String() != want {
		t.Fatalf("Wanted: \"%s\"\nGot: \"%s\"", want, p.String())
	}
}

func TestPosInts(t *testing.T) {
	a, b := 3, 5
	p := &Pos{X: a, Y: b}
	x, y := p.Ints()
	if x != a || y != b {
		t.Fatalf("Wanted: %d, %d\nGot: %d, %d", a, b, x, y)
	}
}

func TestSpriteFace(t *testing.T) {
	s := &Sprite{Tiles: [][]termbox.Cell{{gorogue.DefaultPlayer}}}
	var v SpriteFace
	v = s
	if s != v.GetSprite() {
		t.Fatal("Sprite does not implement SpriteFace.")
	}
}
