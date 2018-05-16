package gorogue

import (
	// "bytes"
	"fmt"
	"log"
	"net"
	// "net/rpc/jsonrpc"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var Log *log.Logger

// CatchSignals runs a goroutine that handles POSIX signals.
// It gets called by NewServer.
func CatchSignals() {
	c := make(chan os.Signal, 2)
	signal.Notify(c)
	go func() {
		for sig := range c {
			log.Println(sig)
			switch sig {
			case syscall.SIGTERM:
				fallthrough
			case syscall.SIGKILL:
				fallthrough
			case syscall.SIGINT:
				log.Println("Exiting...")
				os.Exit(1)
			}
		}
	}()
}

// NewClient initializes a connection to a server. NewClient must be called to
// connect to an online game, but can be ignored if using a local model.
//
// Each process can connect to only one server, meaning each call to NewClient
// will overwrite the previous connection.
func NewClient(c Client, host, port string) error {
	// Disconnect the previous session, if any
	// if stdConn != nil {
	// 	stdConn.Disconnect()
	// 	stdConn = nil
	// }
	stdConn = c
	conn, err := net.Dial("tcp", fmt.Sprintf("%s%s", host, port))
	if err != nil {
		return err
	}
	// defer stdConn.Disconnect()

	var addr []byte = make([]byte, 24)
	if _, err := conn.Read(addr); err != nil {
		return err
	}
	fmt.Println(string(addr), conn.LocalAddr())
	// stdConn.SetAddr(string(bytes.Trim(addr, "\x00")))
	// stdConn.SetRPC(jsonrpc.NewClient(conn))
	if err := stdConn.Init(); err != nil {
		panic(err)
	}
	stdUI.Run()
	return nil
}

// NewServer starts a server on the given port. Servers are used to control
// the World in an online game.
func NewServer(s Server, port string) {
	CatchSignals()
	s.SetPort(port)
	s.HandleRequests()
}

func SetLog(w io.Writer) {
	Log = log.New(w, "", log.LstdFlags)
}
