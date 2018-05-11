package main

import (
	"flag"
	rogue "github.com/sbrow/gorogue"
)

func main() {
	var port = flag.String("port", ":6060", "The port to host from. Must include the colon.")
	flag.Parse()
	_ = rogue.NewServer(*port, rogue.NewMap(100, 100, "Map"))
}
