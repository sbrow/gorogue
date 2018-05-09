package gorogue

import (
	"fmt"
	. "github.com/sbrow/gorogue"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Game logic is handled on the server.
//
// Each world gets at least one goroutine, with each active map getting its own
// goroutine as well.
type Server struct {
	Host string
	Port string
	Maps []*Map
}

func Start(host, port string, maps ...*Map) *Server {
	s := &Server{host, port, maps}
	go s.HandleRequests()
	return s
}

func (s *Server) HandleRequests() {
	server := rpc.NewServer()
	server.Register(s)

	l, err := net.Listen("tcp", fmt.Sprint(s.Port))
	if err != nil {
		panic(err)
	}

	for {
		if conn, err := l.Accept(); err != nil {
			panic(err)
		} else {
			log.Println("Connection established!")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}

func (s *Server) Map(args *string, reply *Map) error {
	for _, m := range s.Maps {
		if m.Name == *args {
			*reply = *m
			break
		}
	}
	return nil
}

func (s *Server) Move(args *MoveArgs, reply *bool) error {
	s.Maps[0].Actors[0].SetPos(args.Points[0])
	*reply = true
	return nil
}

// func (s *Server) Spawn(args *Player, reply *Map) error {
func (s *Server) Spawn(args Actors, reply *SpawnReply) error {
	m := s.Maps[0]
	args[0].SetPos(Point{5, 5})
	m.Actors = append(m.Actors, args...)
	*reply = SpawnReply{&m.Name, m.Actors}
	return nil
}
