package main

import (
	rogue "github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/client"
	server "github.com/sbrow/gorogue/server"
)

func main() {
	host, port := "localhost", ":6061"
	go server.Start(host, port, rogue.NewMap(100, 100, "Map"))
	client.Connect(host, port)

