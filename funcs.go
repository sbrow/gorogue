package gorogue

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
	"runtime"
)

var Log *log.Logger

var stdConn Client

// NewRemoteClient initializes a Client connection to a server. NewClient must be called to
// connect to an online game, but can be ignored if using a local model.
//
// Each process can connect to only one server, meaning each call to NewClient
// will overwrite the previous connection.
func NewRemoteClient(c RemoteClient, host, port string) error {
	// Disconnect the previous session, if any
	if stdConn != nil {
		switch stdConn.(type) {
		case RemoteClient:
			stdConn.(RemoteClient).Disconnect()
		}
		stdConn = nil
	}
	remoteConn := stdConn.(RemoteClient)
	remoteConn = c
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
	remoteConn.Run()
	return nil
}

// NewServer starts a server on the given port. Servers are used to control
// the World in an online game.
func NewServer(s Server, port string) {
	s.SetPort(port)
	s.HandleRequests()
}

func SetLog(name string) (*os.File, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("Something went wrong.")
	}
	f, err := os.Create(filepath.Join(filepath.Dir(filename), name+".log"))
	if err != nil {
		return nil, err
	}
	Log = log.New(f, "", log.LstdFlags)
	return f, nil
}
