package systems

import (
	"reflect"
	"sync"
	"testing"

	"engo.io/ecs"
	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/components"
)

type Scene struct {
	world *ecs.World
	sys   interface {
		ecs.SystemAddByInterfacer
		Sizer
	}
	ent interface {
		components.BasicFace
	}
	face interface{}
	mut  *sync.Mutex
}

var s *Scene

func init() {
	s = &Scene{mut: &sync.Mutex{}}
}

func AddEnt(t *testing.T) {
	s.world.AddEntity(s.ent)
	if s.sys.Size() == 0 {
		t.Error("entity was not added to the system")
	}
}

func RemoveEnt(t *testing.T) {

	n := s.sys.Size()
	s.sys.Remove(*s.ent.GetBasicEntity())
	if s.sys.Size() != n-1 {
		t.Fatal("entity was not removed from the system")
	}
	s.sys.Remove(*s.ent.GetBasicEntity())
	if s.sys.Size() != n-1 {
		t.Error("attempted to remove a non-existent entity from a system")
	}
}

func SetupWorld(t *testing.T) {
	s.world = &ecs.World{}
	s.world.AddSystemInterface(s.sys, s.face, nil)
	if len(s.world.Systems()) != 1 {
		t.Fatal("world could not add system with given interface")
	}
}

func AbleTest(t *testing.T) {
	t.Run("New", SetupWorld)
	t.Run("Add", AddEnt)
	t.Run("Remove", RemoveEnt)
}

func TestAction(t *testing.T) {
	s.mut.Lock()
	s.sys = &Action{}
	var i *Actionable
	s.face = i
	s.ent = &struct {
		ecs.BasicEntity
		components.Action
	}{
		ecs.NewBasic(),
		components.Action{Name: "Test", Caller: "Client"},
	}
	t.Run("Actionable", AbleTest)
	s.mut.Unlock()
}

func TestRender(t *testing.T) {
	s.mut.Lock()
	s.sys = &Render{}
	var i *Renderable
	s.face = i
	s.ent = &struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{X: 0, Y: 0},
		components.Sprite{Tiles: [][]termbox.Cell{{gorogue.DefaultPlayer}}},
	}
	t.Run("New", SetupWorld)
	t.Run("Add", AddEnt)
	t.Run("Update", RenderUpdate)
	t.Run("Remove", RemoveEnt)
	termbox.Close()
	s.mut.Unlock()
}

func RenderUpdate(t *testing.T) {
	r, ok := s.sys.(*Render)
	if !ok {
		t.Fatal("given system is not a Render system")
	}

	r.Update(0)
	data, err := r.PrintScreen()
	if err != nil {
		t.Fatal(err)
	}

	got := string(data)[0:1]
	want := string(gorogue.DefaultPlayer.Ch)
	if got != want {
		t.Errorf("Wanted: \n\"%s\"\nGot: \n\"%s\"", want, got)
	}
}

func TestInput(t *testing.T) {
	w := &ecs.World{}
	w.AddSystem(&Input{})
	sys := w.Systems()[0].(*Input)
	eq := reflect.DeepEqual(sys.binds, gorogue.Keybinds)
	if !eq {
		t.Error("keys were not bound correctly")
	}
	if sys.world != w {
		t.Error("input system was not initialized correctly")
	}
	if sys.player == nil {
		t.Error("input system did not have a player in it")
	}
	// TODO: Fix
	/*
		if sys.player != nil {
			t.Fatal("input system had an unexpected entity in it")
		}
		e := ecs.NewBasic()
		sys.Add(&e)
		if *sys.player.BasicEntity != e {
			t.Error("input system.Add had an unexpected result")
		}
		sys.Remove(e)
	*/
	sys.Remove(*sys.player.GetBasicEntity())
	if sys.player != nil {
		t.Error("input system could not remove the entity")
	}
}

func TestInputUpdate(t *testing.T) {
	w := &ecs.World{}
	w.AddSystem(&Input{})
	var i *Actionable
	w.AddSystemInterface(&Action{}, i, nil)
	w.Update(0)
	sys, ok := w.Systems()[1].(*Action)
	if !ok {
		t.Fatal("err")
	}
	if sys.Size() == 0 {
		t.Error("action was not sent from input system to action system")
	}
}
