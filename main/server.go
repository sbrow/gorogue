package main

import (
	rogue "github.com/sbrow/gorogue"
	server "github.com/sbrow/gorogue/server"
)

func main() {
	host, port := "localhost", ":6061"
	_ = server.Start(host, port, rogue.NewMap(100, 100, "Map"))
}
