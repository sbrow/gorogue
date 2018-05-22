package gorogue

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"path/filepath"
	"runtime"
)

// NewRemoteClient initializes a Client connection to a server.
// NewRemoteClient must be called in order to connect to an online game.
// For local games, use NewClient instead.
//
// Each process can have only one active Client, meaning each call to NewRemoteClient
// or NewClient will overwrite the previous client.
func NewRemoteClient(c RemoteClient, host, port string) error {
	// Disconnect the previous session, if any
	if StdConn != nil {
		switch StdConn.(type) {
		case RemoteClient:
			StdConn.(RemoteClient).Disconnect()
		}
		StdConn = nil
	}
	StdConn = c
	remoteConn := StdConn.(RemoteClient)
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

// SetLog sets the output for the standard Logger.
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

func Dist(a, b Point) float64 {
	x1, y1 := a.Ints()
	x2, y2 := b.Ints()
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}
