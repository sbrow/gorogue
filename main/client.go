package main

import (
	"flag"
	"github.com/sbrow/gorogue/client"
)

func main() {
	var host = flag.String("host", "localhost", "The host to connect to.")
	var port = flag.String("port", ":6061", "The port to host from, must include the colon.")
	flag.Parse()
	client.Connect(*host, *port)
}
