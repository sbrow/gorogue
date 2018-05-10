package main

import (
	"github.com/sbrow/gorogue/client"
)

func main() {
	host, port := "localhost", ":6061"
	client.Connect(host, port)
}
