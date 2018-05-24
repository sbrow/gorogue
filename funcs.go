package gorogue

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// NewRemoteClient initializes a Client connection to a server.
// NewRemoteClient must be called in order to connect to an online game.
// For local games, use NewClient instead.
//
// Each process can have only one active Client, meaning each call to NewRemoteClient
// or NewClient will overwrite the previous client.
func NewRemoteClient(c RemoteClient, host, port string) error {
	// Disconnect the previous session, if any
	if stdConn != nil {
		switch stdConn.(type) {
		case RemoteClient:
			stdConn.(RemoteClient).Disconnect()
		}
		stdConn = nil
	}
	stdConn = c
	remoteConn := stdConn.(RemoteClient)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		return err
	}
	defer remoteConn.Disconnect()

	var addr []byte = make([]byte, 24)
	if _, err := conn.Read(addr); err != nil {
		return err
	}
	fmt.Println(string(addr), conn.LocalAddr())
	remoteConn.SetAddr(string(bytes.Trim(addr, "\x00")))
	remoteConn.SetRPC(jsonrpc.NewClient(conn))
	if err := remoteConn.Init(); err != nil {
		panic(err)
	}
	remoteConn.Init()
	remoteConn.Run()
	return nil
}

// NewServer starts a server on the given port. Servers are used to control
// the World in an online game.
/*func NewServer(s Server, port string) {
	s.SetPort(port)
	s.HandleRequests()
}*/

func HandleAction(a *Action) error {
	if stdConn != nil {
		return stdConn.HandleAction(a)
	}
	return errors.New("No client set.")
}
