package main

import (
	"flag"
	engine "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/example"
	"time"
)

func main() {
	var port = flag.String("port", ":6060", "The port to host from, must include the colon.")
	flag.Parse()
	go engine.NewServer(&example.Server{}, *port)
	err := engine.NewClient(&example.Client{}, "localhost", *port)
	if err != nil {
		panic(err)
	}
}
