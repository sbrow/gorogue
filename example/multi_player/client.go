package main

import (
	"flag"
	"github.com/sbrow/gorogue"
	"github.com/sbrow/gorogue/assets"
)

func main() {
	var host = flag.String("host", "localhost", "The host to connect to.")
	var port = flag.String("port", ":6061", "The port to host from, must include the colon.")
	flag.Parse()
	gorogue.SetLog("client")
	gorogue.NewRemoteClient(&assets.ExampleRemoteClient{}, *host, *port)
}
