package main

import (
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/example"
	. "github.com/sbrow/gorogue/example/single_player"
)

func main() {
	f, err := engine.SetLog("local_game")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := example.NewMap(15, 24, "Map_1")
	w := example.NewWorld(m)
	c := NewClient(w)
	if err := c.Init(); err != nil {
		panic(err)
	}
	c.Run()
}
