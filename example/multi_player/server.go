package main

import (
	"flag"
	"github.com/sbrow/gorogue/assets"
)

func main() {
	var port = flag.String("port", ":6060", "The port to host from. Must include the colon.")
	flag.Parse()
	assets.NewServer(*port)
}
