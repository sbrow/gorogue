package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	termbox "github.com/nsf/termbox-go"
	"github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/components"
)

// Input is a system that turns keys into actions for the Action system.
type Input struct {
	player *playerEnt
	binds  map[gorogue.Key]gorogue.Action
	world  *ecs.World
}

// Add adds an entity to the system.
func (i *Input) Add(ent *ecs.BasicEntity, pos *components.Pos, sprite *components.Sprite) {
	i.player = &playerEnt{ent, pos, sprite}
}

// KeyPressed checks to see what action the given key is bound to.
// It returns a KeyNotBoundError if the key is unbound.
func (i *Input) KeyPressed(key gorogue.Key) (*gorogue.Action, error) {
	if act, ok := i.binds[key]; ok {
		return &act, nil
	}
	return nil, &gorogue.KeyNotBoundError{Key: key}
}

// New Initializes the system and connects it to the world.
//
// Keybinds are pulled from gorogue.Keybinds.
func (i *Input) New(w *ecs.World) {
	i.world = w
	i.binds = gorogue.Keybinds
	gorogue.SetLog("test.log", true)
	if !termbox.IsInit {
		if err := termbox.Init(); err != nil {
			// Panicking is bad style here, but the engo API forces our hand.
			panic(err)
		}
	}

	// TODO: Fix
	p := &struct {
		ecs.BasicEntity
		components.Pos
		components.Sprite
	}{
		ecs.NewBasic(),
		components.Pos{X: 0, Y: 0},
		components.Sprite{Tiles: [][]termbox.Cell{{gorogue.DefaultPlayer}}},
	}
	i.Add(p.GetBasicEntity(), p.GetPos(), p.GetSprite())
	i.world.AddEntity(p)
}

// Poll gets a key from the user and returns the Action it's mapped to.
// If the Key pressed is bound to an Action, the action is returned along with a nil error.
// if the Key pressed isn't bound to an Action, a nil Action is returned along
// with a KeyNotBoundError.
func (i *Input) Poll() (*gorogue.Action, error) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			key := &gorogue.Key{}
			key.Mod = ev.Mod
			if ev.Ch != 0 {
				key.Ch = ev.Ch
			} else {
				key.Key = ev.Key
			}
			return i.KeyPressed(*key)
		}
	}
}

// Remove removes the given entity from the system.
func (i *Input) Remove(e ecs.BasicEntity) {
	i.player = nil
}

// Update gets input from the user and if that input is bound
// to an action, the action is sent to the Action system.
func (i *Input) Update(dt float32) {
	action, err := i.Poll()
	if err != nil {
		gorogue.Log.Println(err)
	} else if action != nil {
		if action.Name == "Quit" {
			engo.Exit()
		}
		i.world.AddEntity(&struct {
			ecs.BasicEntity
			components.Action
		}{
			ecs.NewBasic(),
			components.Action{Name: action.Name, Caller: action.Caller},
		})
	}
}

// player is the player character.
type playerEnt struct {
	*ecs.BasicEntity
	*components.Pos
	*components.Sprite
}
