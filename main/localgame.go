package main

import (
	"flag"
	rogue "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/client"
)

func main() {
	var port = flag.String("port", ":6060", "The port to host from, must include the colon.")
	flag.Parse()
	go rogue.NewServer(port, rogue.NewMap(100, 100, "Map"))
	client.Connect("localhost", port)
}
