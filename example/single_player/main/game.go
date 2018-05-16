package main

import (
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/example"
	. "github.com/sbrow/gorogue/example/single_player"
	"log"
	"os"
)

func main() {
	f, err := os.Create("game.log")
	if err != nil {
		log.Panic("Shit.")
	}
	defer f.Close()
	engine.SetLog(f)

	m := example.NewMap(15, 24, "Map_1")
	w := example.NewWorld(m)
	c := NewClient(w)
	if err := c.Init(); err != nil {
		panic(err)
	}
	c.Run()
}
