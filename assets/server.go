package assets

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ExampleServer struct {
	*ExampleWorld
	port string
	// conns []string
}

func NewServer(port string) {
	e := &ExampleServer{ExampleWorld: NewWorld()}
	e.NewMap(5, 5)
	e.port = port
	// e.conns = []string{}
	e.HandleRequests()
}

/*func (e *ExampleServer) Conns() []string {
	return e.conns
}
*/
func (e *ExampleServer) Disconnect(args *string, reply *string) error {
	Log.Printf("%s Disconnected.\n", *args)
	actor := e.players[*args]
	actor.Map().Remove(*args)
	/*	for i, conn := range e.conns {
			if conn == *args {
				e.conns = append(e.conns[:i], e.conns[i+1:]...)
			}
		}
	*/return nil
}

/*func (e *ExampleServer) HandleAction(a *Action, reply *string) error {
	fmt.Println("Conns:", e.conns)
	return e.ExampleWorld.HandleAction(a, reply)
}
*/
func (e *ExampleServer) HandleRequests() {
	server := rpc.NewServer()
	server.RegisterName("Server", e)

	l, err := net.Listen("tcp", fmt.Sprint(e.port))
	if err != nil {
		panic(err)
	}
	Log.Println("Server started on port:", e.port)
	Log.Println("Waiting for players...")
	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			Log.Printf("%s Connected.", conn.RemoteAddr())
			addr := fmt.Sprint(conn.RemoteAddr())
			conn.Write([]byte(addr))
			// e.conns = append(e.conns, addr)
			e.Spawn(NewAction("Spawn", addr, NewPlayer("Player")))
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

func (e *ExampleServer) Map(args *int, reply *[][]Tile) error {
	*reply = e.Maps()[*args].AllTiles()
	return nil
}

func (e *ExampleServer) Port() string {
	return e.port
}

func (e *ExampleServer) SetPort(port string) {
	e.port = port
}
