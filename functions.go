package gorogue

import (
	"bytes"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// NewClient initializes a connection to a server. NewClient must be called to
// connect to an online game, but can be ignored if using a local model.
//
// Each process can connect to only one server, meaning each call to NewClient
// will overwrite the previous connection.
func NewClient(c Client, host, port string) error {
	// Disconnect the previous session, if any
	if stdConn != nil {
		stdConn.Disconnect()
		stdConn = nil
	}
	stdConn = c
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		return err
	}
	defer stdConn.Disconnect()

	var addr []byte = make([]byte, 24)
	if _, err := conn.Read(addr); err != nil {
		return err
	}
	fmt.Println(string(addr), conn.LocalAddr())
	stdConn.SetAddr(string(bytes.Trim(addr, "\x00")))
	stdConn.SetRPC(jsonrpc.NewClient(conn))
	stdUI = stdConn.Init()
	// stdUI.Run()
	return nil
}

// NewServer starts a server on the given port.
func NewServer(s Server, port string) {
	CatchSignals()
	s.SetPort(port)
	s.HandleRequests()
}
