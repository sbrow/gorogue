package gorogue

import (
	"bytes"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// Connect initializes a connection to a server. It must be called before all other
// functions.
func NewClient(c Client, host, port string) error {
	stdConn = c
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		return err
	}

	defer conn.Close()
	var addr []byte = make([]byte, 24)
	if _, err := conn.Read(addr); err != nil {
		return err
	}
	stdConn.SetAddr(string(bytes.Trim(addr, "\x00")))
	stdConn.SetRPC(jsonrpc.NewClient(conn))
	stdUI = stdConn.Init()
	stdUI.Run()
	return nil
}

func NewServer(s Server, port string) {
	// server := reflect.TypeOf((*Server)(nil)).Elem()
	// vType := reflect.TypeOf(v)
	// if !reflect.TypeOf(v).Implements(server) {
	// log.Panicf("Interface type %s does not implement %s\n", vType, server)
	// }
	// srv := v.(Server) //v.(Server)
	CatchSignals()
	s.SetPort(port)
	s.HandleRequests()
}
