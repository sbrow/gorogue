package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"flag"
	"log"
)

type ServerScene struct{}

func (*ServerScene) Preload() {}

func (*ServerScene) Setup(u engo.Updater) {
	w, ok := u.(*ecs.World)
	if !ok {
		panic("Updater is not a world!")
	}
	log.Printf("World created: %+v\n", w)
}

func (*ServerScene) Type() string { return "DedicatedServer" }

func main() {
	var port = flag.String("port", ":6060", "The port to host from. Must include the colon.")
	flag.Parse()
	opts := engo.RunOptions{
		HeadlessMode: true,
		FPSLimit:     20,
	}
	log.Println(*port)
	scene := &ServerScene{}
	engo.Run(opts, scene)
}
