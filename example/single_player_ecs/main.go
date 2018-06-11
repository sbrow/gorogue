package main

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/sbrow/gorogue/systems"
)

type singlePlayerScene struct{}

func (s *singlePlayerScene) Opts() engo.RunOptions {
	return engo.RunOptions{
		HeadlessMode: true,
		FPSLimit:     20,
	}
}

func (s *singlePlayerScene) Preload() {}

func (s *singlePlayerScene) Setup(u engo.Updater) {
	w, ok := u.(*ecs.World)
	if !ok {
		panic("Updater is not a world!")
	}
	log.Printf("World created: %+v\n", w)

	var r *systems.Renderable
	w.AddSystemInterface(&systems.Render{}, r, nil)
	w.AddSystem(&systems.Input{})
}

func (*singlePlayerScene) Type() string { return "SinglePlayerGame" }

func main() {
	scene := &singlePlayerScene{}
	engo.Run(scene.Opts(), scene)
}
