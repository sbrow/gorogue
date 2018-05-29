package systems

import (
	"fmt"

	"engo.io/ecs"
	"github.com/sbrow/gorogue/components"
)

// Action handles actions and sends them to the appropriate systems.
type Action struct {
	ents  []actionEnt
	world *ecs.World
}

// Add adds the action to the system.
func (a *Action) Add(ent *ecs.BasicEntity, action *components.Action) {
	a.ents = append(a.ents, actionEnt{ent, action})
}

// AddByInterface adds an entity to the system, asserting that it implements Actionable.
// It will panic if the assertion fails.
func (a *Action) AddByInterface(i ecs.Identifier) {
	obj, ok := i.(Actionable)
	if !ok {
		panic(fmt.Sprintf("%s is not Actionable.", i))
	}
	a.Add(obj.GetBasicEntity(), obj.GetAction())
}

// New Initializes the system and connects it to the world.
func (a *Action) New(w *ecs.World) {
	a.world = w
	a.ents = []actionEnt{}
}

// Update handles all queued actions.
func (a *Action) Update(dt float32) {
	// TODO: Implement
}

// Remove removes the entity from the system.
func (a *Action) Remove(e ecs.BasicEntity) {
	i := -1
	for j, e := range a.ents {
		if e.ID() == e.ID() {
			i = j
			break
		}
	}
	if i >= 0 {
		a.ents = append(a.ents[:i], a.ents[i+1:]...)
	}
}

// Size returns how many entities are in the system.
func (a *Action) Size() int {
	return len(a.ents)
}

// ActionEnt is an unexported struct for Actionable entities.
type actionEnt struct {
	*ecs.BasicEntity
	*components.Action
}

// Actionable is the interface that must be filled for an entity to be added
// to an Action system.
type Actionable interface {
	components.BasicFace
	components.ActionFace
}
