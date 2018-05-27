package systems

import (
	"errors"
	"testing"

	"engo.io/ecs"

	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/components"
)

func worldSetup() (*ecs.World, error) {
	w := &ecs.World{}
	var renderable *Renderable
	w.AddSystemInterface(&Render{}, renderable, nil)
	if len(w.Systems()) != 1 {
		return nil, errors.New("World could not add Render system with Renderable interface.")
	}
	return w, nil
}

func TestRenderable(t *testing.T) {
	w, err := worldSetup()
	if err != nil {
		t.Fatal(err)
	}
	e := struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{0, 0},
		components.Sprite{[][]termbox.Cell{[]termbox.Cell{gorogue.DefaultPlayer}}},
	}
	w.AddEntity(&e)
	r, ok := w.Systems()[0].(*Render)
	if !ok {
		t.Fatal("World system[0] is not a RenderSystem.")
	}
	if len(r.ents) == 0 {
		t.Fatal("Entity was not added to the system.")
	}
}

func TestRender_Remove(t *testing.T) {
	w, err := worldSetup()
	if err != nil {
		t.Fatal(err)
	}
	e := struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{0, 0},
		components.Sprite{[][]termbox.Cell{[]termbox.Cell{gorogue.DefaultPlayer}}},
	}
	w.AddEntity(&e)

	r, ok := w.Systems()[0].(*Render)
	if !ok {
		t.Fatal("World system[0] is not a RenderSystem.")
	}
	n := len(r.ents)
	r.Remove(e.BasicEntity)
	if len(r.ents) != n-1 {
		t.Fatal("Entity was not removed from RenderSystem.")
	}
	r.Remove(e.BasicEntity)
	if len(r.ents) != n-1 {
		t.Fatal("Non-existant Entity was removed from RenderSystem.")
	}
}

func TestRender_Update(t *testing.T) {
	w, err := worldSetup()
	if err != nil {
		t.Fatal(err)
	}
	e := struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{0, 0},
		components.Sprite{[][]termbox.Cell{[]termbox.Cell{gorogue.DefaultPlayer}}},
	}
	w.AddEntity(&e)
	r, ok := w.Systems()[0].(*Render)
	if !ok {
		t.Fatal("World system[0] is not a RenderSystem.")
	}
	r.Update(0)
	data, err := r.PrintScreen()
	if err != nil {
		t.Fatal(err)
	}
	termbox.Close()

	got := string(data)[0:1]
	want := string(gorogue.DefaultPlayer.Ch)
	if got != want {
		t.Fatalf("Wanted: \n\"%s\"\nGot: \n\"%s\"", want, got)
	}
}
